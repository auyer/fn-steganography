# fn-steganography

![OpenFaaS](https://camo.githubusercontent.com/e400c2b9b42deb6d444a3a509ccdba416f76fe2d/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6f70656e666161732d253343332d626c75652e737667)
[![Go Report Card](https://goreportcard.com/badge/github.com/auyer/steganography)](https://goreportcard.com/report/github.com/auyer/steganography)
[![LICENSE MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)
[![CircleCI](https://circleci.com/gh/auyer/fn-steganography.svg?style=svg)](https://circleci.com/gh/auyer/fn-steganography)
[![codecov](https://codecov.io/gh/auyer/fn-steganography/branch/master/graph/badge.svg)](https://codecov.io/gh/auyer/fn-steganography)

Run LSB steganography enconding/decoding on OpenFaaS! 


| Original              | Encoded           |
| --------------------  | ------------------|
| <img src="https://github.com/auyer/steganography/raw/master/examples/stegosaurus.png"/>        | <img src="https://github.com/auyer/steganography/raw/master/examples/encoded_stegosaurus.png"/>

The second image contains the first paragaph of the description of a stegosaurus on [Wikipidia](https://en.wikipedia.org/wiki/Stegosaurus), also available in [examples/message.txt](https://raw.githubusercontent.com/auyer/steganography/master/examples/message.txt) as an example.


## How to use it

This function recieves a JSON with 3 fields:
- An image encoded in base 64;
- Boolean argument indicating encoding/decoding mode;
- A message (in case of encoding).

-------
## Encoding:
Example on a smaller pure white image:
```json
{	
    "message" : "hello stego" ,
    "image" : "iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAYAAADED76LAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAFUlEQVQY02P8DwQMeAATAwEwPBQAABtuBAy91jkOAAAAAElFTkSuQmCC",
    "encode" : true
}
``` 

Will produce as a result:

```json
{
    "image": "iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAIAAABLbSncAAAATUlEQVR4nFSOURbEIAwCh33e/8hhX8S2xh/NEMBVVQASBiHO+W2at1u7hebeTJ+ycklZ8HBkdCdpOHoO89t3OqxZ8ETlm9c68A8AAP//+vgUFEMX0moAAAAASUVORK5CYIKJUE5HDQoaCgAAAA1JSERSAAAACAAAAAgIAgAAAEttKdwAAABOSURBVHicVI1BEsAgDAIXx/8/WTqKadqLkgSWudYCkDDoyP0wWmJcMoc92Mfh2jPzSTH0YYSRzId0EnvOzm/f7bD+BYVCl98kngAAAP//+5QUFaju17IAAAAASUVORK5CYII="
}
```

______
## Decoding

```json
{
    "encode": false,
    "image": "iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAIAAABLbSncAAAATUlEQVR4nFSOURbEIAwCh33e/8hhX8S2xh/NEMBVVQASBiHO+W2at1u7hebeTJ+ycklZ8HBkdCdpOHoO89t3OqxZ8ETlm9c68A8AAP//+vgUFEMX0moAAAAASUVORK5CYIKJUE5HDQoaCgAAAA1JSERSAAAACAAAAAgIAgAAAEttKdwAAABOSURBVHicVI1BEsAgDAIXx/8/WTqKadqLkgSWudYCkDDoyP0wWmJcMoc92Mfh2jPzSTH0YYSRzId0EnvOzm/f7bD+BYVCl98kngAAAP//+5QUFaju17IAAAAASUVORK5CYII="
}
```

Will produce as a result:

```json
{
    "message": "hello stego"
}
```
____

### Encode and Call function with curl:
```bash
(echo -n '{"encode: true, "message": "hello stego","image": "'; base64 ~/path_to_pic.png; echo '"}') |
curl -H "Content-Type: application/json" -d @-  http://127.0.0.1:8080/function/steganography
```

-----
### Attributions
 - Stegosaurus Picture By Matt Martyniuk - Own work, CC BY-SA 3.0, https://commons.wikimedia.org/w/index.php?curid=42215661
