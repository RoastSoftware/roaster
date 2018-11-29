Security
========
This document states which security issues that has been taken into consideration when developing this site.

TLS
---
The connection from the client to the web server is made over TLS with a Amazon signed certificate.
This is implemented on the Elastic Loadbalancer of Elastic Beanstalk.
The request then propagates through our VPC (Virtual private cloud) in AWS as plain HTTP. 

CSRF
----
The site has implemented protection against CSRF (Cross Site Request Forgery), 
> this is a manin the middle attack that abuses the browser’s automatic cookie submission for cross-origin requests to issue state changing requests on the user’s behalf. 
> In other words, the attack is meant to trick users into issuing requests by abusing browser session cookie management.
[1](https://blog.securityevaluators.com/cracking-javas-rng-for-csrf-ea9cacd231d2)

In order to increase the security against CSRSF attacks we have chosen to use a rolling random token scheme. This means that the base senario for protection against CSRF attacks gives you a session wide CSRF-cookie and which generates a session wide CSRF-token.
The enhanced implementation still gives you a session wide CSRF-cookie which is used to generate _per request_ CSRF-tokens.

Secure cookies
--------------
Roast.software cookies is stored in the cookie store in the browser. These are validate with the HMAC mechanism and encrypted with AES256.

Password storage
----------------
There's alot of challenges keeping user credentials secured. Our solution is based on well known algorithms, Bcrypt and SHA512. The implementation is made of the official libraries for go, so that we don't make the mistake of "rolling our own crypto". The methodology used is to first hash the password with SHA512, and then strongly hashing it with Bcrypt.

Below quote is found in `/model/user.go`.
> The plaintext password that is provided will first be transformed to a
> hash sum with SHA-512. This is due to that Bcrypt limits the input to 72
> bytes. By hashing the password with SHA-512 more entropy of the original
> password is kept. Also, some implementations of Bcrypt that allows for longer
> passwords can be vulnerable to DoS attacks[0].
>
> The SHA-512 hash sum is then hashed again using Bcrypt. This is because
> SHA-512 is a _fast_ hash algorithm not made for password hashing. Bcrypt is
> designed to be slow and hard to speed up using hardware such as FPGAs and
> ASICs. The work factor is set to 12 which should make the expensive Blowfish
> setup take >250 ms (364.815906 ms precisely on my sucky laptop).
>
> Dropbox has a great article[1] on their password hashing scheme which our
> scheme shares many similarities with, we do not, however, use AES-256 with a
> global pepper (shared global encryption key) which is overkill for our use
> case. They also encode their SHA-512 with base 64, which is not needed in our
> case.
>
> Some implementations of Bcrypt uses a null byte (\x00) to determine the end
> of the input, the Go implementations does _not_ have this problem. So there
> is no need to encode the data as base 64. This is verified using the program
> in cmd/testbcrypt.
>
> The approach used by Dropbox where they encode with base 64 will generate a
> ~88 byte long key, which is then truncated to 72 bytes. This results in a
> input with 64^72 possible combinations. Our approach of not encoding the
> input as base 64 results in a 64 byte long key (the output size of SHA-512),
> where there is 256 possible combinations _per byte_. This results in a input
> with 256^64 possible combinations. Therefore our approach allows for far more
> possible entropy because 64^72 << 256^64 (like 99.999... % more entropy ;)).
>
> [0]: https://arstechnica.com/information-technology/2013/09/long-passwords-are-good-but-too-much-length-can-be-bad-for-security/
> [1]: https://blogs.dropbox.com/tech/2016/09/how-dropbox-securely-stores-your-passwords/


