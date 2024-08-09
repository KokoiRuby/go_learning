## [Protobuf](https://protobuf.dev/)

与语言/平台无关可扩展的**对结构化数据 XML/JSON 进行序列化/反序列化机制** but Smaller/Faster/Simpler。

Most commonly-used data format at Google, extensively in inter-server comm as well as archive storage on disk.

- [gRPC](https://grpc.io/)
- [Google Cloud](https://cloud.google.com/)
- [Envoy Proxy](https://www.envoyproxy.io/)

:smile:

- Compact data storage
- Fast parsing
- Availability in many programming languages
- Optimized functionality through automatically-generated classes

![Compilation workflow showing the creation of a proto file, generated code, and compiled classes](https://protobuf.dev/images/protocol-buffers-concepts.png)

`.proto` 文件定义消息 **Messages** 和 服务 **Services**。

```protobuf
message Person {
	optional string name = 1;
	optional int32 id = 2;
	optional string email = 3;
}
```

Proto Compiler 对 `.proto` 进行构建生成适配各种语言的源码。每个生成的类都有消息字段的 setter 以及序列化/反序列化的方法。

```java
Person john = Person.newBuilder()
    .setId(1234)
    .setName("John Doe")
    .setEmail("jdoe@example.com")
    .build();
output = new FileOutputStream(args[0]);
john.writeTo(output);
```

反序列化使用不同的语言, like C++。

```c++
Person john;
fstream input(argv[1], ios::in | ios::binary);
john.ParseFromIstream(&input);
int id = john.id();
std::string name = john.name();
std::string email = john.email();
```

### [Install](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)

**Compiler**

```bash
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v27.3/protoc-27.3-linux-x86_64.zip
$ extract protoc-27.3-linux-x86_64.zip
$ ln -s /path/to/protoc-27.3-linux-x86_64/bin/protoc /usr/local/bin/protoc
```

**Runtime**

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ ln -s $GOPATH/bin/protoc-gen-go /usr/local/bin
```

### [Example](https://github.com/protocolbuffers/protobuf/tree/main/examples)

[Bazel](https://bazel.build/install/ubuntu)

```bash
$ sudo apt install apt-transport-https curl gnupg -y
$ curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel-archive-keyring.gpg
$ sudo mv bazel-archive-keyring.gpg /usr/share/keyrings
$ echo "deb [arch=amd64 signed-by=/usr/share/keyrings/bazel-archive-keyring.gpg] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
```

```bash
$ sudo apt update && sudo apt install bazel
$ bazel --version
```

[Build](https://github.com/protocolbuffers/protobuf/tree/main/examples#go) by make

```bash
# modify examples/go/go.mod
# go 1.14 → go 1.22
$ make go
```

Run

```bash
# to add a person to the protocol buffer encoded file
./add_person_go addressbook.data
# to view the data
./list_people_go addressbook.data
```

### [Go](https://protobuf.dev/getting-started/gotutorial/) 

A very simple “address book” application that can read and write people’s contact details to and from a file.

Each **person** in the address book has a name, an ID, an email address, and a contact phone number.

1. `.proto` ++ **messages** for each struct u want to secrialize, then specify a name and a type for each field. A **message** is just an aggregate containing a set of typed fields. For more info, pls refer to [Guide](https://protobuf.dev/programming-guides/proto3).

```protobuf
option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";

// If a field is repeated, the field may be repeated any number of times
message AddressBook {
  repeated Person people = 1;
}

// msg def
// The " = 1", " = 2" markers on each element identify the unique “tag” that field uses in the binary encoding.
// Tag numbers 1-15 require one less byte to encode than higher numbers
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;

  google.protobuf.Timestamp last_updated = 5;
}

enum PhoneType {
  PHONE_TYPE_UNSPECIFIED = 0;
  PHONE_TYPE_MOBILE = 1;
  PHONE_TYPE_HOME = 2\;
  PHONE_TYPE_WORK = 3;
}
```

2. Compile to generate the classes to R/W **messages**.

```bash
$ protoc -I=./protobuf_examples/ --go_out=. ./protobuf_examples/addressbook.proto
```

3. Dep

```bash
$ go mod init tutorialpb
$ go mod tidy
```

4. ProtoBuf API

- An `AddressBook` structure with a `People` field.
- A `Person` structure with fields for `Name`, `Id`, `Email` and `Phones`.
- A `Person_PhoneNumber` structure, with fields for `Number` and `Type`.
- The type `Person_PhoneType` and a value defined for each value in the `Person.PhoneType` enum.
- See More in [Guide](https://protobuf.dev/reference/go/go-generated).

5. Write a Message by `Marshal` to serialize

```go
book := &pb.AddressBook{}

// Write the new address book back to disk.
out, err := proto.Marshal(book)
if err != nil {
    log.Fatalln("Failed to encode address book:", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
    log.Fatalln("Failed to write address book:", err)
}
```

6. Read a Message by `Unmarshal`

```go
// Read the existing address book.
in, err := ioutil.ReadFile(fname)
if err != nil {
    log.Fatalln("Error reading file:", err)
}
book := &pb.AddressBook{}
if err := proto.Unmarshal(in, book); err != nil {
    log.Fatalln("Failed to parse address book:", err)
}
```







