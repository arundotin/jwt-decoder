## JWT Decoder CLI

There's nothing fancy in this CLI. It just takes in the accessToken as a flag and prints the decoded `header` & `payload`

## CLI Execution

The access-token needs to be given as a flag to the command line. Sample below

```shell script
./jwtdecoder --accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

Once this command is executed, the header & payload will be printed like below

```

------------------------------
Token Header
------------------------------
{
	"alg": "HS256",
	"typ": "JWT"
}
------------------------------
------------------------------
Token payload
------------------------------
{
	"sub": "1234567890",
	"name": "John Doe",
	"iat": 1516239022
}

```