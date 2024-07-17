### [stdblib](https://pkg.go.dev/std)

|                                                              |                                                              |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)archive |                                                              |
| [tar](https://pkg.go.dev/archive/tar@go1.22.5)               | Package tar implements access to tar archives.               |
| [zip](https://pkg.go.dev/archive/zip@go1.22.5)               | Package zip provides support for reading and writing ZIP archives. |
| [bufio](https://pkg.go.dev/bufio@go1.22.5)                   | Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O. |
| [builtin](https://pkg.go.dev/builtin@go1.22.5)               | Package builtin provides documentation for Go's predeclared identifiers. |
| [bytes](https://pkg.go.dev/bytes@go1.22.5)                   | Package bytes implements functions for the manipulation of byte slices. |
| [cmp](https://pkg.go.dev/cmp@go1.22.5)                       | Package cmp provides types and functions related to comparing ordered values. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)compress |                                                              |
| [bzip2](https://pkg.go.dev/compress/bzip2@go1.22.5)          | Package bzip2 implements bzip2 decompression.                |
| [flate](https://pkg.go.dev/compress/flate@go1.22.5)          | Package flate implements the DEFLATE compressed data format, described in RFC 1951. |
| [gzip](https://pkg.go.dev/compress/gzip@go1.22.5)            | Package gzip implements reading and writing of gzip format compressed files, as specified in RFC 1952. |
| [lzw](https://pkg.go.dev/compress/lzw@go1.22.5)              | Package lzw implements the Lempel-Ziv-Welch compressed data format, described in T. A. Welch, “A Technique for High-Performance Data Compression”, Computer, 17(6) (June 1984), pp 8-19. |
| [zlib](https://pkg.go.dev/compress/zlib@go1.22.5)            | Package zlib implements reading and writing of zlib format compressed data, as specified in RFC 1950. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)container |                                                              |
| [heap](https://pkg.go.dev/container/heap@go1.22.5)           | Package heap provides heap operations for any type that implements heap.Interface. |
| [list](https://pkg.go.dev/container/list@go1.22.5)           | Package list implements a doubly linked list.                |
| [ring](https://pkg.go.dev/container/ring@go1.22.5)           | Package ring implements operations on circular lists.        |
| [context](https://pkg.go.dev/context@go1.22.5)               | Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[crypto](https://pkg.go.dev/crypto@go1.22.5) | Package crypto collects common cryptographic constants.      |
| [aes](https://pkg.go.dev/crypto/aes@go1.22.5)                | Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197. |
| [cipher](https://pkg.go.dev/crypto/cipher@go1.22.5)          | Package cipher implements standard block cipher modes that can be wrapped around low-level block cipher implementations. |
| [des](https://pkg.go.dev/crypto/des@go1.22.5)                | Package des implements the Data Encryption Standard (DES) and the Triple Data Encryption Algorithm (TDEA) as defined in U.S. Federal Information Processing Standards Publication 46-3. |
| [dsa](https://pkg.go.dev/crypto/dsa@go1.22.5)                | Package dsa implements the Digital Signature Algorithm, as defined in FIPS 186-3. |
| [ecdh](https://pkg.go.dev/crypto/ecdh@go1.22.5)              | Package ecdh implements Elliptic Curve Diffie-Hellman over NIST curves and Curve25519. |
| [ecdsa](https://pkg.go.dev/crypto/ecdsa@go1.22.5)            | Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as defined in FIPS 186-4 and SEC 1, Version 2.0. |
| [ed25519](https://pkg.go.dev/crypto/ed25519@go1.22.5)        | Package ed25519 implements the Ed25519 signature algorithm.  |
| [elliptic](https://pkg.go.dev/crypto/elliptic@go1.22.5)      | Package elliptic implements the standard NIST P-224, P-256, P-384, and P-521 elliptic curves over prime fields. |
| [hmac](https://pkg.go.dev/crypto/hmac@go1.22.5)              | Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards Publication 198. |
| [md5](https://pkg.go.dev/crypto/md5@go1.22.5)                | Package md5 implements the MD5 hash algorithm as defined in RFC 1321. |
| [rand](https://pkg.go.dev/crypto/rand@go1.22.5)              | Package rand implements a cryptographically secure random number generator. |
| [rc4](https://pkg.go.dev/crypto/rc4@go1.22.5)                | Package rc4 implements RC4 encryption, as defined in Bruce Schneier's Applied Cryptography. |
| [rsa](https://pkg.go.dev/crypto/rsa@go1.22.5)                | Package rsa implements RSA encryption as specified in PKCS #1 and RFC 8017. |
| [sha1](https://pkg.go.dev/crypto/sha1@go1.22.5)              | Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174. |
| [sha256](https://pkg.go.dev/crypto/sha256@go1.22.5)          | Package sha256 implements the SHA224 and SHA256 hash algorithms as defined in FIPS 180-4. |
| [sha512](https://pkg.go.dev/crypto/sha512@go1.22.5)          | Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256 hash algorithms as defined in FIPS 180-4. |
| [subtle](https://pkg.go.dev/crypto/subtle@go1.22.5)          | Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly. |
| [tls](https://pkg.go.dev/crypto/tls@go1.22.5)                | Package tls partially implements TLS 1.2, as specified in RFC 5246, and TLS 1.3, as specified in RFC 8446. |
| [x509](https://pkg.go.dev/crypto/x509@go1.22.5)              | Package x509 implements a subset of the X.509 standard.      |
| [x509/pkix](https://pkg.go.dev/crypto/x509/pkix@go1.22.5)    | Package pkix contains shared, low level structures used for ASN.1 parsing and serialization of X.509 certificates, CRL and OCSP. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)database |                                                              |
| [sql](https://pkg.go.dev/database/sql@go1.22.5)              | Package sql provides a generic interface around SQL (or SQL-like) databases. |
| [sql/driver](https://pkg.go.dev/database/sql/driver@go1.22.5) | Package driver defines interfaces to be implemented by database drivers as used by package sql. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)debug |                                                              |
| [buildinfo](https://pkg.go.dev/debug/buildinfo@go1.22.5)     | Package buildinfo provides access to information embedded in a Go binary about how it was built. |
| [dwarf](https://pkg.go.dev/debug/dwarf@go1.22.5)             | Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf. |
| [elf](https://pkg.go.dev/debug/elf@go1.22.5)                 | Package elf implements access to ELF object files.           |
| [gosym](https://pkg.go.dev/debug/gosym@go1.22.5)             | Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers. |
| [macho](https://pkg.go.dev/debug/macho@go1.22.5)             | Package macho implements access to Mach-O object files.      |
| [pe](https://pkg.go.dev/debug/pe@go1.22.5)                   | Package pe implements access to PE (Microsoft Windows Portable Executable) files. |
| [plan9obj](https://pkg.go.dev/debug/plan9obj@go1.22.5)       | Package plan9obj implements access to Plan 9 a.out object files. |
| [embed](https://pkg.go.dev/embed@go1.22.5)                   | Package embed provides access to files embedded in the running Go program. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[encoding](https://pkg.go.dev/encoding@go1.22.5) | Package encoding defines interfaces shared by other packages that convert data to and from byte-level and textual representations. |
| [ascii85](https://pkg.go.dev/encoding/ascii85@go1.22.5)      | Package ascii85 implements the ascii85 data encoding as used in the btoa tool and Adobe's PostScript and PDF document formats. |
| [asn1](https://pkg.go.dev/encoding/asn1@go1.22.5)            | Package asn1 implements parsing of DER-encoded ASN.1 data structures, as defined in ITU-T Rec X.690. |
| [base32](https://pkg.go.dev/encoding/base32@go1.22.5)        | Package base32 implements base32 encoding as specified by RFC 4648. |
| [base64](https://pkg.go.dev/encoding/base64@go1.22.5)        | Package base64 implements base64 encoding as specified by RFC 4648. |
| [binary](https://pkg.go.dev/encoding/binary@go1.22.5)        | Package binary implements simple translation between numbers and byte sequences and encoding and decoding of varints. |
| [csv](https://pkg.go.dev/encoding/csv@go1.22.5)              | Package csv reads and writes comma-separated values (CSV) files. |
| [gob](https://pkg.go.dev/encoding/gob@go1.22.5)              | Package gob manages streams of gobs - binary values exchanged between an [Encoder] (transmitter) and a [Decoder] (receiver). |
| [hex](https://pkg.go.dev/encoding/hex@go1.22.5)              | Package hex implements hexadecimal encoding and decoding.    |
| [json](https://pkg.go.dev/encoding/json@go1.22.5)            | Package json implements encoding and decoding of JSON as defined in RFC 7159. |
| [pem](https://pkg.go.dev/encoding/pem@go1.22.5)              | Package pem implements the PEM data encoding, which originated in Privacy Enhanced Mail. |
| [xml](https://pkg.go.dev/encoding/xml@go1.22.5)              | Package xml implements a simple XML 1.0 parser that understands XML name spaces. |
| [errors](https://pkg.go.dev/errors@go1.22.5)                 | Package errors implements functions to manipulate errors.    |
| [expvar](https://pkg.go.dev/expvar@go1.22.5)                 | Package expvar provides a standardized interface to public variables, such as operation counters in servers. |
| [flag](https://pkg.go.dev/flag@go1.22.5)                     | Package flag implements command-line flag parsing.           |
| [fmt](https://pkg.go.dev/fmt@go1.22.5)                       | Package fmt implements formatted I/O with functions analogous to C's printf and scanf. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)go |                                                              |
| [ast](https://pkg.go.dev/go/ast@go1.22.5)                    | Package ast declares the types used to represent syntax trees for Go packages. |
| [build](https://pkg.go.dev/go/build@go1.22.5)                | Package build gathers information about Go packages.         |
| [build/constraint](https://pkg.go.dev/go/build/constraint@go1.22.5) | Package constraint implements parsing and evaluation of build constraint lines. |
| [constant](https://pkg.go.dev/go/constant@go1.22.5)          | Package constant implements Values representing untyped Go constants and their corresponding operations. |
| [doc](https://pkg.go.dev/go/doc@go1.22.5)                    | Package doc extracts source code documentation from a Go AST. |
| [doc/comment](https://pkg.go.dev/go/doc/comment@go1.22.5)    | Package comment implements parsing and reformatting of Go doc comments, (documentation comments), which are comments that immediately precede a top-level declaration of a package, const, func, type, or var. |
| [format](https://pkg.go.dev/go/format@go1.22.5)              | Package format implements standard formatting of Go source.  |
| [importer](https://pkg.go.dev/go/importer@go1.22.5)          | Package importer provides access to export data importers.   |
| [parser](https://pkg.go.dev/go/parser@go1.22.5)              | Package parser implements a parser for Go source files.      |
| [printer](https://pkg.go.dev/go/printer@go1.22.5)            | Package printer implements printing of AST nodes.            |
| [scanner](https://pkg.go.dev/go/scanner@go1.22.5)            | Package scanner implements a scanner for Go source text.     |
| [token](https://pkg.go.dev/go/token@go1.22.5)                | Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates). |
| [types](https://pkg.go.dev/go/types@go1.22.5)                | Package types declares the data types and implements the algorithms for type-checking of Go packages. |
| [version](https://pkg.go.dev/go/version@go1.22.5)            | Package version provides operations on [Go versions] in [Go toolchain name syntax]: strings like "go1.20", "go1.21.0", "go1.22rc2", and "go1.23.4-bigcorp". |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[hash](https://pkg.go.dev/hash@go1.22.5) | Package hash provides interfaces for hash functions.         |
| [adler32](https://pkg.go.dev/hash/adler32@go1.22.5)          | Package adler32 implements the Adler-32 checksum.            |
| [crc32](https://pkg.go.dev/hash/crc32@go1.22.5)              | Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32, checksum. |
| [crc64](https://pkg.go.dev/hash/crc64@go1.22.5)              | Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. |
| [fnv](https://pkg.go.dev/hash/fnv@go1.22.5)                  | Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. |
| [maphash](https://pkg.go.dev/hash/maphash@go1.22.5)          | Package maphash provides hash functions on byte sequences.   |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[html](https://pkg.go.dev/html@go1.22.5) | Package html provides functions for escaping and unescaping HTML text. |
| [template](https://pkg.go.dev/html/template@go1.22.5)        | Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[image](https://pkg.go.dev/image@go1.22.5) | Package image implements a basic 2-D image library.          |
| [color](https://pkg.go.dev/image/color@go1.22.5)             | Package color implements a basic color library.              |
| [color/palette](https://pkg.go.dev/image/color/palette@go1.22.5) | Package palette provides standard color palettes.            |
| [draw](https://pkg.go.dev/image/draw@go1.22.5)               | Package draw provides image composition functions.           |
| [gif](https://pkg.go.dev/image/gif@go1.22.5)                 | Package gif implements a GIF image decoder and encoder.      |
| [jpeg](https://pkg.go.dev/image/jpeg@go1.22.5)               | Package jpeg implements a JPEG image decoder and encoder.    |
| [png](https://pkg.go.dev/image/png@go1.22.5)                 | Package png implements a PNG image decoder and encoder.      |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)index |                                                              |
| [suffixarray](https://pkg.go.dev/index/suffixarray@go1.22.5) | Package suffixarray implements substring search in logarithmic time using an in-memory suffix array. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[io](https://pkg.go.dev/io@go1.22.5) | Package io provides basic interfaces to I/O primitives.      |
| [fs](https://pkg.go.dev/io/fs@go1.22.5)                      | Package fs defines basic interfaces to a file system.        |
| [ioutil](https://pkg.go.dev/io/ioutil@go1.22.5)              | Package ioutil implements some I/O utility functions.        |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[log](https://pkg.go.dev/log@go1.22.5) | Package log implements a simple logging package.             |
| [slog](https://pkg.go.dev/log/slog@go1.22.5)                 | Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs. |
| [syslog](https://pkg.go.dev/log/syslog@go1.22.5)             | Package syslog provides a simple interface to the system log service. |
| [maps](https://pkg.go.dev/maps@go1.22.5)                     | Package maps defines various functions useful with maps of any type. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[math](https://pkg.go.dev/math@go1.22.5) | Package math provides basic constants and mathematical functions. |
| [big](https://pkg.go.dev/math/big@go1.22.5)                  | Package big implements arbitrary-precision arithmetic (big numbers). |
| [bits](https://pkg.go.dev/math/bits@go1.22.5)                | Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types. |
| [cmplx](https://pkg.go.dev/math/cmplx@go1.22.5)              | Package cmplx provides basic constants and mathematical functions for complex numbers. |
| [rand](https://pkg.go.dev/math/rand@go1.22.5)                | Package rand implements pseudo-random number generators suitable for tasks such as simulation, but it should not be used for security-sensitive work. |
| [rand/v2](https://pkg.go.dev/math/rand/v2@go1.22.5)          | Package rand implements pseudo-random number generators suitable for tasks such as simulation, but it should not be used for security-sensitive work. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[mime](https://pkg.go.dev/mime@go1.22.5) | Package mime implements parts of the MIME spec.              |
| [multipart](https://pkg.go.dev/mime/multipart@go1.22.5)      | Package multipart implements MIME multipart parsing, as defined in RFC 2046. |
| [quotedprintable](https://pkg.go.dev/mime/quotedprintable@go1.22.5) | Package quotedprintable implements quoted-printable encoding as specified by RFC 2045. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[net](https://pkg.go.dev/net@go1.22.5) | Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. |
| [http](https://pkg.go.dev/net/http@go1.22.5)                 | Package http provides HTTP client and server implementations. |
| [http/cgi](https://pkg.go.dev/net/http/cgi@go1.22.5)         | Package cgi implements CGI (Common Gateway Interface) as specified in RFC 3875. |
| [http/cookiejar](https://pkg.go.dev/net/http/cookiejar@go1.22.5) | Package cookiejar implements an in-memory RFC 6265-compliant http.CookieJar. |
| [http/fcgi](https://pkg.go.dev/net/http/fcgi@go1.22.5)       | Package fcgi implements the FastCGI protocol.                |
| [http/httptest](https://pkg.go.dev/net/http/httptest@go1.22.5) | Package httptest provides utilities for HTTP testing.        |
| [http/httptrace](https://pkg.go.dev/net/http/httptrace@go1.22.5) | Package httptrace provides mechanisms to trace the events within HTTP client requests. |
| [http/httputil](https://pkg.go.dev/net/http/httputil@go1.22.5) | Package httputil provides HTTP utility functions, complementing the more common ones in the net/http package. |
| [http/pprof](https://pkg.go.dev/net/http/pprof@go1.22.5)     | Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool. |
| [mail](https://pkg.go.dev/net/mail@go1.22.5)                 | Package mail implements parsing of mail messages.            |
| [netip](https://pkg.go.dev/net/netip@go1.22.5)               | Package netip defines an IP address type that's a small value type. |
| [rpc](https://pkg.go.dev/net/rpc@go1.22.5)                   | Package rpc provides access to the exported methods of an object across a network or other I/O connection. |
| [rpc/jsonrpc](https://pkg.go.dev/net/rpc/jsonrpc@go1.22.5)   | Package jsonrpc implements a JSON-RPC 1.0 ClientCodec and ServerCodec for the rpc package. |
| [smtp](https://pkg.go.dev/net/smtp@go1.22.5)                 | Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321. |
| [textproto](https://pkg.go.dev/net/textproto@go1.22.5)       | Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP. |
| [url](https://pkg.go.dev/net/url@go1.22.5)                   | Package url parses URLs and implements query escaping.       |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[os](https://pkg.go.dev/os@go1.22.5) | Package os provides a platform-independent interface to operating system functionality. |
| [exec](https://pkg.go.dev/os/exec@go1.22.5)                  | Package exec runs external commands.                         |
| [signal](https://pkg.go.dev/os/signal@go1.22.5)              | Package signal implements access to incoming signals.        |
| [user](https://pkg.go.dev/os/user@go1.22.5)                  | Package user allows user account lookups by name or id.      |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[path](https://pkg.go.dev/path@go1.22.5) | Package path implements utility routines for manipulating slash-separated paths. |
| [filepath](https://pkg.go.dev/path/filepath@go1.22.5)        | Package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths. |
| [plugin](https://pkg.go.dev/plugin@go1.22.5)                 | Package plugin implements loading and symbol resolution of Go plugins. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[reflect](https://pkg.go.dev/reflect@go1.22.5) | Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[regexp](https://pkg.go.dev/regexp@go1.22.5) | Package regexp implements regular expression search.         |
| [syntax](https://pkg.go.dev/regexp/syntax@go1.22.5)          | Package syntax parses regular expressions into parse trees and compiles parse trees into programs. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[runtime](https://pkg.go.dev/runtime@go1.22.5) | Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines. |
| [cgo](https://pkg.go.dev/runtime/cgo@go1.22.5)               | Package cgo contains runtime support for code generated by the cgo tool. |
| [coverage](https://pkg.go.dev/runtime/coverage@go1.22.5)     |                                                              |
| [debug](https://pkg.go.dev/runtime/debug@go1.22.5)           | Package debug contains facilities for programs to debug themselves while they are running. |
| [metrics](https://pkg.go.dev/runtime/metrics@go1.22.5)       | Package metrics provides a stable interface to access implementation-defined metrics exported by the Go runtime. |
| [pprof](https://pkg.go.dev/runtime/pprof@go1.22.5)           | Package pprof writes runtime profiling data in the format expected by the pprof visualization tool. |
| [race](https://pkg.go.dev/runtime/race@go1.22.5)             | Package race implements data race detection logic.           |
| [trace](https://pkg.go.dev/runtime/trace@go1.22.5)           | Package trace contains facilities for programs to generate traces for the Go execution tracer. |
| [slices](https://pkg.go.dev/slices@go1.22.5)                 | Package slices defines various functions useful with slices of any type. |
| [sort](https://pkg.go.dev/sort@go1.22.5)                     | Package sort provides primitives for sorting slices and user-defined collections. |
| [strconv](https://pkg.go.dev/strconv@go1.22.5)               | Package strconv implements conversions to and from string representations of basic data types. |
| [strings](https://pkg.go.dev/strings@go1.22.5)               | Package strings implements simple functions to manipulate UTF-8 encoded strings. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[sync](https://pkg.go.dev/sync@go1.22.5) | Package sync provides basic synchronization primitives such as mutual exclusion locks. |
| [atomic](https://pkg.go.dev/sync/atomic@go1.22.5)            | Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[syscall](https://pkg.go.dev/syscall@go1.22.5) | Package syscall contains an interface to the low-level operating system primitives. |
| [js](https://pkg.go.dev/syscall/js@go1.22.5)                 | Package js gives access to the WebAssembly host environment when using the js/wasm architecture. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[testing](https://pkg.go.dev/testing@go1.22.5) | Package testing provides support for automated testing of Go packages. |
| [fstest](https://pkg.go.dev/testing/fstest@go1.22.5)         | Package fstest implements support for testing implementations and users of file systems. |
| [iotest](https://pkg.go.dev/testing/iotest@go1.22.5)         | Package iotest implements Readers and Writers useful mainly for testing. |
| [quick](https://pkg.go.dev/testing/quick@go1.22.5)           | Package quick implements utility functions to help with black box testing. |
| [slogtest](https://pkg.go.dev/testing/slogtest@go1.22.5)     | Package slogtest implements support for testing implementations of log/slog.Handler. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)text |                                                              |
| [scanner](https://pkg.go.dev/text/scanner@go1.22.5)          | Package scanner provides a scanner and tokenizer for UTF-8-encoded text. |
| [tabwriter](https://pkg.go.dev/text/tabwriter@go1.22.5)      | Package tabwriter implements a write filter (tabwriter.Writer) that translates tabbed columns in input into properly aligned text. |
| [template](https://pkg.go.dev/text/template@go1.22.5)        | Package template implements data-driven templates for generating textual output. |
| [template/parse](https://pkg.go.dev/text/template/parse@go1.22.5) | Package parse builds parse trees for templates as defined by text/template and html/template. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[time](https://pkg.go.dev/time@go1.22.5) | Package time provides functionality for measuring and displaying time. |
| [tzdata](https://pkg.go.dev/time/tzdata@go1.22.5)            | Package tzdata provides an embedded copy of the timezone database. |
| ![img](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg)[unicode](https://pkg.go.dev/unicode@go1.22.5) | Package unicode provides data and functions to test some properties of Unicode code points. |
| [utf16](https://pkg.go.dev/unicode/utf16@go1.22.5)           | Package utf16 implements encoding and decoding of UTF-16 sequences. |
| [utf8](https://pkg.go.dev/unicode/utf8@go1.22.5)             | Package utf8 implements functions and constants to support text encoded in UTF-8. |
| [unsafe](https://pkg.go.dev/unsafe@go1.22.5)                 | Package unsafe contains operations that step around the type safety of Go programs. |