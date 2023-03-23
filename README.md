# Golang Part Of Speech Tagger
Golang (GO) Parts Of Speech Tagger (POST)
Hence: GOalPOST

# Purpose
This tool will perform Part of Speech Tagging on a block of text. 
As an example the sentence:
`The quick brown fox jumps over the lazy dog who worked for the corrupt CDC and FDA.`
is processed by the tool to identify all the Nouns, Proper nouns, Verbs, even
organizations, emojis, people, and email addresses. It then outputs the following:
```
The ;  DT ;  O  <---- DT == Determinant
quick ;  JJ ;  O <---- JJ == Adjective
brown ;  NN ;  O <---- NN == Noun
fox ;  NN ;  O <---- NN == Noun
jumps ;  VBZ ;  <---- VBZ == Verb, Third Person Singular
over ;  IN ;  O <--- IN == Conjunction or Preposition
...
FDA ;  GPE    <---- GPE == Geographical or Political Entity.
```
This is a smaller part of a broader project which originally used 
Python NLTK, but I wanted it as a standalone, cross-platform, highly portable tool. 

This tool uses Penn-Treebank format (https://en.wikipedia.org/wiki/Treebank)
A list of tags used to identify the parts of speech is in a Key at the bottom of this document. 

# Use Cases
- Extract names, proper-nouns, et al from large swaths of text.
- Find common names/proper-nouns across many files (especially when used with tools likehttps://github.com/s7ephen/seacrane/wiki#ocr-images-in-a-directory)
- With extracted text from PDFs, Ebooks, audio-files, images, the ability to find commonwords based on their parts of speech is powerful. This is meant to facilitate this.
 
# Installation
This tool does not require any other libraries. It was designed to be highly
portable and cross-platform. Precompiled binaries are available in [bin/](./bin/)

# Building
Precompiled binares are in [bin/](./bin/) but if you wanna build it from in the
directory just do:
`go build goalpost`
or 
`make`
or if you want to make all the distribution:
`make all`

there is also `make windows`, `make x86` (linux), `make osx`, `make arm-osx`, and othersjust have a look at the Makefile for all of them.

# Supported Platforms
|OS Arch| Supported? | Download Link|
|:-:|:-:|-| 
|Windows x86| 🗸 | [bin/goalpost.exe](./bin/goalpost.exe) | 
|Linux x86| 🗸 | [bin/goalpost](./bin/goalpost) |
|Linux ARM| 🗸 | [bin/goalpost-arm](./bin/goalpost-arm) |
|Linux MIPS| 🗸 | [bin/goalpost-mips](./bin/goalpost-mips) |
|Linux MIPS (le)| 🗸 | [bin/goalpost-mips-le](./bin/goalpost-mips-le) |
|Linux MIPS (sf)| 🗸 | [bin/goalpost-mips-sf](./bin/goalpost-mips-sf) |
|OSX ARM| 🗸 | [bin/goalpost-osxarm](./bin/goalpost-osxarm) |
|OSX x86| 🗸 | [bin/goalpost-osx](./bin/goalpost-osx) |
|Android ARM| 🗸 | [bin/goalpost-android](./bin/goalpost-android) |
|Android APK bundle| 🗸 | (coming) | 
|Apple iOS (arm)| built but not on App Store | NA | 
|JS/WebAssembly| 𐄂 | NA | 

# Example Usage:
This tool will print human-readable form when it recieves input via pipe. 
```
$ ./bin/goalpost --help        
Usage of ./bin/goalpost:
  -f string
    	File to perform Part-Of-Speech-Tagging (PoST) on. (default "./file.txt")
$ cat test.txt | ./bin/goalpost
	 [+] STDIN input found, using it instead of file argument.
************************************************************
******************* TOKENIZED/TAGGED TEXT ******************
************************************************************
--- Format is: Text; Tag; Label ---
The ;  DT ;  O
quick ;  JJ ;  O
brown ;  NN ;  O
...
```
It will write the output to a file when the `-f` option is used.
```
$ ./bin/goalpost -f ./example_text.txt 
	 [+] No input found on STDIN, using filename argument instead
		 [+] filename: ./example_text.txt
		 [+] Running PoST on: example_text.txt
	 [+] Writing PoST to:  ./example_text.txt.goalpost_json-1918462
	 [+] Tagged this many items:  134
$
```
# Developer Notes:
Unlike NLTK this tool is meant to be portable. One side-effect is that it does not use a corpus kept on the local disk which helps it to eliminate false-positives. So there will be some. But what is gained in portability/ease-of-use far outweighs this.

# Output File Format:
When `goalpost` receives file as input it does not print anything to the screen
it will simply write the analyzed output json file to disk.
Example:
```
$ ./goalpost -f test.txt      
	 [+] No input found on STDIN, using filename argument instead
		 [+] filename: test.txt
		 [+] Running PoST on: test.txt

         [+] Writing PoST to:  ./test.txt.goalpost_json-2324552945
	 [+] Tagged this many items:  19
$ 
```
The contents of `./test.txt.goalpost_json-2324552945` is 
```
{
    "GPE": [
        {
            "Label": "GPE",
            "Tag": "",
            "Text": "FDA"
        }
    ],
    "Text": "The quick brown fox jumps over the lazy dog who worked for the corrupt CDC and FDA.",
    "Tokens": [
        {
            "Label": "O",
            "Tag": "DT",
            "Text": "The"
        },
        {
            "Label": "O",
            "Tag": "JJ",
            "Text": "quick"
        },
        {
            "Label": "O",
            "Tag": "NN",
            "Text": "brown"
        },
        {
            "Label": "O",
            "Tag": "NN",
            "Text": "fox"
        },
        {
            "Label": "O",
            "Tag": "VBZ",
            "Text": "jumps"
        },
        {
            "Label": "O",
            "Tag": "IN",
            "Text": "over"
        },
        {
            "Label": "O",
            "Tag": "DT",
            "Text": "the"
        },
        {
            "Label": "O",
            "Tag": "JJ",
            "Text": "lazy"
        },
        {
            "Label": "O",
            "Tag": "NN",
            "Text": "dog"
        },
        {
            "Label": "O",
            "Tag": "WP",
            "Text": "who"
        },
        {
            "Label": "O",
            "Tag": "VBD",
            "Text": "worked"
        },
        {
            "Label": "O",
            "Tag": "IN",
            "Text": "for"
        },
        {
            "Label": "O",
            "Tag": "DT",
            "Text": "the"
        },
        {
            "Label": "O",
            "Tag": "JJ",
            "Text": "corrupt"
        },
        {
            "Label": "O",
            "Tag": "NNP",
            "Text": "CDC"
        },
        {
            "Label": "O",
            "Tag": "CC",
            "Text": "and"
        },
        {
            "Label": "B-GPE",
            "Tag": "NNP",
            "Text": "FDA"
        },
        {
            "Label": "O",
            "Tag": ".",
            "Text": "."
        }
    ]
}
```
# Token/Tag Key:

| TAG        | DESCRIPTION                               |
|------------|-------------------------------------------|
| `(`        | left round bracket                        |
| `)`        | right round bracket                       |
| `,`        | comma                                     |
| `:`        | colon                                     |
| `.`        | period                                    |
| `''`       | closing quotation mark                    |
| ``` `` ``` | opening quotation mark                    |
| `#`        | number sign                               |
| `$`        | currency                                  |
| `CC`       | conjunction, coordinating                 |
| `CD`       | cardinal number                           |
| `DT`       | determiner                                |   
| `EX`       | existential there                         |   
| `FW`       | foreign word                              |   
| `IN`       | conjunction, subordinating or preposition |
| `JJ`       | adjective                                 |
| `JJR`      | adjective, comparative                    |
| `JJS`      | adjective, superlative                    |
| `LS`       | list item marker                          |
| `MD`       | verb, modal auxiliary                     |
| `NN`       | noun, singular or mass                    |
| `NNP`      | noun, proper singular                     |
| `NNPS`     | noun, proper plural                       |
| `NNS`      | noun, plural                              |
| `PDT`      | predeterminer                             |
| `POS`      | possessive ending                         |
| `PRP`      | pronoun, personal                         |
| `PRP$`     | pronoun, possessive                       |   
| `RB`       | adverb                                    |   
| `RBR`      | adverb, comparative                       |   
| `RBS`      | adverb, superlative                       |
| `RP`       | adverb, particle                          |
| `SYM`      | symbol                                    |
| `TO`       | infinitival to                            |
| `UH`       | interjection                              |
| `VB`       | verb, base form                           |
| `VBD`      | verb, past tense                          |
| `VBG`      | verb, gerund or present participle        |
| `VBN`      | verb, past participle                     |
| `VBP`      | verb, non-3rd person singular present     |   
| `VBZ`      | verb, 3rd person singular present         |
| `WDT`      | wh-determiner                             |
| `WP`       | wh-pronoun, personal                      |
| `WP$`      | wh-pronoun, possessive                    |   
| `WRB`      | wh-adverb                                 |   


