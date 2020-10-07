import "@babel/polyfill";
import {Terminal} from "xterm";
import {FitAddon} from "xterm-addon-fit";

Swal.fire({
    title: 'Please enter your WNAL host and port',
    input: 'text',
    inputAttributes: {
        autocapitalize: 'off'
    },
    showCancelButton: true,
    confirmButtonText: 'connect',
    showLoaderOnConfirm: true,
    allowOutsideClick: false
}).then((result) => {
    if (result.value) {

        let term = new Terminal();

        const fitAddon = new FitAddon();
        term.loadAddon(fitAddon);

        term.open(document.getElementById('terminal'));

        fitAddon.fit();

        let host = "ws://" + result.value + "/status";

        let socket = new WebSocket(host)

        socket.onerror = function () {
            term.writeln("WNAL: FAILED TO CONNECT TO " + host)
        }

        socket.onclose = function () {
            term.writeln("WNAL: connection closed")
        }

        socket.onmessage = function (data) {
            let m = JSON.parse(data.data).message
            if (/\n/.test(m)) {
                let t = m.split("\n")
                for (let i = 0; i < t.length; i++) {
                    if (i == (t.length - 1)) {
                        console.log("last")
                        term.write(t[i])
                    } else {
                        term.writeln(t[i])
                    }
                }

            } else {
                term.write(m)
            }
        }

        term.onData( (data) => {
            socket.send(data)
        });
    }
})