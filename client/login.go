package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/WilliamOdinson/instant-messaging/common/message"
)

func login(userId int, userPwd string) error {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8889")

	if err != nil {
		fmt.Println("net.Dial err =", err)
		return err
	}
	defer conn.Close()

	// Prepare the login message
	var loginMes = message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}

	// Serialize loginMes
	loginData, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
	}

	// Create a Message instance with type LoginMesType and data as the serialized loginMes
	var mes = message.Message{
		Type: message.LoginMesType,
		Data: string(loginData),
	}

	// Serialize mes
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return err
	}
	// Send the serialized message to the server
	// First send the length of the message
	// First, obtain the length of mesData and convert it into a byte slice.
	var pkgLen = uint32(len(mesData))
	var buf [4]byte // use array because length is fixed (4 bytes for uint32) and lighter than slice
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err =", err)
		return err
	}
	fmt.Printf("Client Length of the message %d\ncontent: %s", len(mesData), mesData)
	fmt.Println(binary.BigEndian.Uint32(buf[0:4]))
	// _, err = conn.Write(mesData)
	// if err != nil {
	// fmt.Println("conn.Write err =", err)
	// return err
	// }

	return nil
}
