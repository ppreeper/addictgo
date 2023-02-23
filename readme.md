//The Nitty Gritty
//Passing Secrets

//You can pass the AD details at runtime:

//addict --url ldaps://[address] --user [user]@[domain] --pass [pass]

//As environmental variables:

//export ADDICT_URL=ldaps://[address]
//export ADDICT_USER=[user]@[domain]
//export ADDICT_PASS=[pass]

//Or in ./config.json:

//git clone https://github.com/dthree/addict.git
//cd addict
//vim ./config.json

//{
//...
//"user": "[user]@[domain]",
//"pass": "[pass]",
//"url": "ldaps://[address]"
//}

//Authentication

//This service defaults to no authentication. I can't and won't try to guess your flavor.

//Addict uses express. The file ./middleware.js at the root of the directory exposes the app so you can add middleware hooks for auth logic.
//LDAP vs LDAPS

//If you connect to Active Directory over plain LDAP, it will refuse certain write operations including adding a user and changing a password. To make things even better, Windows Server doesn't support LDAPS out of the box. You're going to have to set up the Domain Controller as a cert authority by installing the Active Directory Certificate Services Role.

//Here's a good tutorial on that.
