e std::io::net::tcp::TcpListener;
use std::io::net::ip::{Ipv4Addr, SocketAddr};
use std::io::{Acceptor, Listener};

let addr = SocketAddr { ip: Ipv4Addr(127, 0, 0, 1), port: 80 };
let listener = TcpListener::bind(addr);

// bind the listener to the specified address
let mut acceptor = listener.listen();

// accept connections and process them
for stream in acceptor.incoming() {
    spawn(proc() {
        handle_client(stream);
    });
}

// close the socket server
drop(acceptor);
