package main 
import ( 
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
	"github.com/skip2/go-qrcode"
	"os"
	"net"
 	"bufio"
	"strings"
)
// fucntion which generates the qr code 
// input para usename which is string and a secert and retuen 
// string andn error
func generateQR(username string, secret string) (string, error) {
    timestamp := time.Now().Unix()
	// takes a time from the system 
    data := fmt.Sprintf("%s:%d:%s", username, timestamp, secret)
	// formates the data into username and tine and a secert
    hash := sha256.Sum256([]byte(data))
	// make a  unique hash from that 
    uniqueData := hex.EncodeToString(hash[:])
	// convert that hash into a string ( encode )
    // make  a qr code using the lib using this new string hash 
    qr, err := qrcode.New(uniqueData, qrcode.Medium)
    if err != nil {
        return "", err
    }
    // save the qr code with the username
    filename := fmt.Sprintf("%s_qr.png", username)
    err = qr.WriteFile(256, filename)
    if err != nil {
        return "", err
    }
    return filename, nil
}
func validateQR(scannedData string) bool {
    // In reality, decode and check against a stored hash or database.
    // For simplicity, check if it's a valid hex string (from our hash).
    _, err := hex.DecodeString(scannedData)
    return err == nil && len(scannedData) == 64 // SHA-256 hex length
}

func firewallServer() {
	// setups a lister on port 8080 
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Printf("Error starting server: %v\n", err)
        return
    }
	// defer from cleaning up 
    defer listener.Close()

    fmt.Println("Firewall listening on :8080")
    // for ever loop ( actively listeing without any break )
    for {
        conn, err := listener.Accept()
		/// for acceping 
        if err != nil {
            continue
        }
		// using a go rotuine to handel connection without breaking the connection of 
		// lister
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
	// defer for clean up 
    reader := bufio.NewReader(conn)
    scannedData, _ := reader.ReadString('\n')
    scannedData = strings.TrimSpace(scannedData)
    
    if validateQR(scannedData) {
        fmt.Fprintln(conn, "Access granted to network!")
    } else {
        fmt.Fprintln(conn, "Invalid QR code. Access denied.")
    }
}

func main() {
     if len(os.Args) > 1 && os.Args[1] == "server" {
        firewallServer()
        return
    }
    username := os.Args[1]
    secret := os.Args[2]
    
    filename, err := generateQR(username, secret)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("QR code generated: %s\nScan this to request access.\n", filename)
}