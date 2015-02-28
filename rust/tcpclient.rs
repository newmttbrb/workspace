use std::io::net::ip::SocketAddr;
use std::io::net::tcp::TcpStream;

let addr = from_str::<SocketAddr>("127.0.0.1:8080").unwrap();
let mut socket = TcpStream::connect(addr).unwrap();
socket.write(bytes!("GET / HTTP/1.0\n\n"));
let response = socket.read_to_end();
