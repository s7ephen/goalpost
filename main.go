package main

import (
    "github.com/jdkato/prose/v2"
    "fmt"
    "log"
    "encoding/json"
    "flag"
    "bufio"
    "os"
    "io"
    "io/ioutil"
    "path/filepath"
    "crypto/md5"
)

var pln = fmt.Println

func demo_tokens_entities_sentences() {
    // Create a new document with the default configuration:
    doc, err := prose.NewDocument("Go is an open-source programming language created at Google.")
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over the doc's tokens:
    for _, tok := range doc.Tokens() {
        fmt.Println(tok.Text, tok.Tag, tok.Label)
        // Go NNP B-GPE
        // is VBZ O
        // an DT O
        // ...
    }
    fmt.Println("\n==========\n")

    // Iterate over the doc's named-entities:
    for _, ent := range doc.Entities() {
        fmt.Println(ent.Text, ent.Label)
        // Go GPE
        // Google GPE
    }
    fmt.Println("\n==========\n")

    // Iterate over the doc's sentences:
    for _, sent := range doc.Sentences() {
        fmt.Println(sent.Text)
        // Go is an open-source programming language created at Google.
    }
    fmt.Println("\n==========\n")
}

func do_tokens(idoc *[]byte) {
    pln("************************************************************")
    pln("******************* TOKENIZED/TAGGED TEXT ******************")
    pln("************************************************************")
    doc, err := prose.NewDocument(string(*idoc))
    if err != nil {
        log.Fatal(err)
    }
    // Iterate over the doc's tokens:
    fmt.Println ("--- Format is: Text; Tag; Label ---")
    for _, tok := range doc.Tokens() {
        fmt.Println(tok.Text,"; ", tok.Tag, "; ", tok.Label)
    }
/*
    pln("************************************************************")
    pln("*********** JSON VERSION OF TOKENIZED/TAGGED TEXT *********")
    pln("************************************************************")
    u, err := json.Marshal(doc.Tokens())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(u))
*/
}

func do_entities(idoc *[]byte) {
    pln("************************************************************")
    pln("******* PERSONS, GEOGRAPHICAL, and POLITICAL ENTITIES ******")
    pln("************************************************************")
    doc, err := prose.NewDocument(string(*idoc))
    if err != nil {
        log.Fatal(err)
    }
    // Iterate over the doc's named-entities:
    fmt.Println ("--- Format is: Text; Label ---")
    for _, ent := range doc.Entities() {
        fmt.Println(ent.Text, "; ", ent.Label)
        // Go GPE
        // Google GPE
    }
/*
    pln("************************************************************")
    pln("*** JSON VERSION PERSONS, GEOGRAPHICAL/POLITICAL ENTITIES **")
    pln("************************************************************")
    u, err := json.Marshal(doc.Entities())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(u))
*/
}

type GoalpostJSON struct {
    FullFile string
    MD5Digest string
    FileName string
    Text string
    GPE []TokenizedBlock
    Tokens []TokenizedBlock
}

type TokenizedBlock struct{
    Text string
    Tag string
    Label string
}

var JSONOutput GoalpostJSON

func md5file (file *string) string {
    var digest string
    f, err := os.Open(*file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    hash := md5.New()
    _, err = io.Copy(hash,f)
    if err !=nil {
        log.Fatal(err)
    }
    digest = fmt.Sprintf("%x", hash.Sum(nil))
    return digest
}
func post_file(file *string){
    var fname string
    var tagcount uint64 //I dunno why I am making this so big, whatever.
    tagcount = 0
    fname = filepath.Clean(*file)
    pln("\t\t [+] Running PoST on:", fname)
    var text []byte
    fileContent, err :=os.ReadFile(fname)
    if err != nil {
        log.Fatal(err)
    }
    text=fileContent
    doc, err := prose.NewDocument(string(text))
    if err != nil {
        log.Fatal(err)
    }
    tempFile, err := ioutil.TempFile("./", fname+".goalpost_json-")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()
    fmt.Println("\t [+] Writing PoST to: ", tempFile.Name())
// *********  BUILDING OUR JSON OUTPUT BELOW ***********
    JSONOutput.Text=string(doc.Text)
    JSONOutput.FullFile=*file
    JSONOutput.FileName=fname
    JSONOutput.MD5Digest=md5file(file)
    var tempBlock TokenizedBlock
    for _, ent := range doc.Entities() {
	tempBlock.Text = ent.Text
	tempBlock.Tag = ""
	tempBlock.Label = ent.Label
        JSONOutput.GPE = append(JSONOutput.GPE, tempBlock)
	tagcount+=1
    }
    for _, tok := range doc.Tokens() {
	tempBlock.Text = tok.Text
	tempBlock.Tag = tok.Tag
	tempBlock.Label = tok.Label
        JSONOutput.Tokens = append(JSONOutput.Tokens,tempBlock)
	tagcount+=1
    }
// ************* DONE BUILDING OUTPUT ******************
    fmt.Println("\t [+] Tagged this many items: ", tagcount)
    json.NewEncoder(tempFile).Encode(JSONOutput)
}

func main() {
    //demo_tokens_entities_sentences()
    var argsvar string
    flag.StringVar(&argsvar,"f", "./file.txt", "File to perform Part-Of-Speech-Tagging (PoST) on.")
    flag.Parse()
    // check if there is somethinig to read on STDIN
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        var stdin []byte
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            stdin = append(stdin, scanner.Bytes()...)
        }
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
	pln("\t [+] STDIN input found, using it instead of file argument.")
        //fmt.Printf("STDIN = %s\n", stdin)
	do_tokens(&stdin)
	do_entities(&stdin)
	pln("\t [+] Remember, to get this output in JSON format use '-f' option instead.")
    } else {
	pln("\t [+] No input found on STDIN, using filename argument instead")
        pln("\t\t [+] filename:", argsvar)
        post_file(&argsvar)
    }
}
