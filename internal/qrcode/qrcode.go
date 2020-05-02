package qrcode

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
    "net"
	
	"github.com/skip2/go-qrcode"
)

func Load(r *mux.Router){
	r.HandleFunc("/api/v1/qrcode", GetQrCodePNG).Methods("GET")
}

func GetQrCodePNG(w http.ResponseWriter, r *http.Request) {
	msg := "http://" + GetOutboundIP().String() + ":3000"
	qrCodeImage := GenerateQrCodePNG(msg)
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(200)
	w.Write(qrCodeImage)
}

func GenerateQrCodePNG(msg string) ([]byte){
	var png []byte
	png, err := qrcode.Encode(msg, qrcode.Medium, 256)
    if err != nil {
        log.Fatal(err)
    }
	return png
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}