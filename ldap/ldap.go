package ldap

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/ppreeper/str"
)

// GetScope uses bind username to determine basic domain scope
func GetScope(bindUser string) (scope string) {
	domain := strings.Split(strings.Split(bindUser, "@")[1], ".")
	for _, v := range domain {
		scope += fmt.Sprintf("dc=%s,", v)
	}
	return str.LJustLen(scope, len(scope)-1)
}

// LDAP configuration stuct
type LDAP struct {
	URL      string
	Username string
	Password string
	Scope    string
	Port     string
	Conn     *ldap.Conn
}

func (a *LDAP) HostPort() string {
	return ":" + a.Port
}

// Dial connect LDAP server
func (a *LDAP) Dial() (*ldap.Conn, error) {
	l, err := ldap.DialURL(a.URL)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return l, err
}

// Bind to LDAP server
func (a *LDAP) Bind(conn *ldap.Conn, user, pass string) error {
	err := conn.Bind(user, pass)
	return err
}

// Authenticate against AD server
func (a *LDAP) Authenticate(user, pass string) map[string]any {
	d := make(map[string]any)
	d["data"] = false
	conn, err := a.Dial()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.Bind(a.UserDN(user), pass)
	if err != nil {
		d["data"] = false
	} else {
		d["data"] = true
	}
	return d
}

// UserDN finds dn from sAMAccountName
func (a *LDAP) UserDN(user string) (dn string) {
	args := new(LDAPArgs)
	args.Fields = "distinguishedName"
	filter := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", user)
	sr := a.Search(filter, *args)
	if len(sr) != 0 {
		return fmt.Sprintf("%s", sr["distinguishedName"])
	}
	return
}

type LDAPArgs struct {
	Fields string `query:"fields"`
	EQ     string `query:"eq"`
	NE     string `query:"ne"`
	LT     string `query:"lt"`
	GT     string `query:"gt"`
	GTE    string `query:"gte"`
	LTE    string `query:"lte"`
	Like   string `query:"like"`
	Sort   string `query:"sort"`
	Order  string `query:"order"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Start  int    `query:"start"`
	End    int    `query:"end"`
	Q      string `query:"q"`
}

// Search search ldap based on filter and attributes
func (a *LDAP) Search(filter string, args LDAPArgs) map[string]any {
	var attributes []string
	if len(args.Fields) == 0 {
		attributes = append(attributes, "*")
	} else {
		attributes = strings.Split(args.Fields, ",")
	}
	// fmt.Println("fields: ", args.Fields, "attributes: ", attributes, len(attributes))
	// fmt.Println("q: ", args.Q)
	// fmt.Println("page: ", args.Page, "start: ", args.Start, "end: ", args.End, "limit: ", args.Limit)
	// fmt.Println("sort: ", args.Sort, " order: ", args.Order)

	conn, err := a.Dial()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer conn.Close()
	err = conn.Bind(a.Username, a.Password)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	searchRequest := ldap.NewSearchRequest(
		a.Scope,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		filter, attributes, nil,
	)
	sr, err := conn.Search(searchRequest)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	start := 0
	if args.Start > 0 && args.Start < len(sr.Entries) {
		start = args.Start
	}
	end := len(sr.Entries)
	if args.End > 0 && args.End < len(sr.Entries) && args.End > start {
		end = args.End
	}
	if args.Limit > 0 && args.Limit <= (len(sr.Entries)-start) {
		end = start + args.Limit
	}

	// fmt.Println("start: ", start, "end: ", end, "count: ", len(sr.Entries))

	d := make(map[string]any)
	var dd []map[string]any
	for k, e := range sr.Entries {
		if k >= start {
			if k < end {
				// fmt.Println(e)
				// e.PrettyPrint(2)
				// fmt.Println(len(e.Attributes))
				if len(e.Attributes) > 0 {
					m := make(map[string]any)
					for _, attr := range e.Attributes {
						if len(attr.Values) == 1 {
							m[attr.Name] = attr.Values[0]
							// dd = append(dd, m)
						} else if len(attr.Values) > 1 {
							m[attr.Name] = attr.Values
							// dd = append(dd, m)
						}
					}
					dd = append(dd, m)
				}
			}
		}
	}
	d["data"] = dd
	return d
}

type Attrs []ldap.Attribute

func (a *LDAP) AddRecord(dn string, attributes Attrs) error {
	conn, err := a.Dial()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer conn.Close()
	err = conn.Bind(a.Username, a.Password)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	dn += "," + a.Scope
	addRequest := ldap.AddRequest{
		DN:         dn,
		Attributes: attributes,
		Controls:   []ldap.Control{},
	}
	err = conn.Add(&addRequest)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return err
}
