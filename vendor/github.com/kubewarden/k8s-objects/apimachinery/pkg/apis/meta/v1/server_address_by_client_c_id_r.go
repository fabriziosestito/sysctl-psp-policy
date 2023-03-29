// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ServerAddressByClientCIDR ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.
//
// swagger:model ServerAddressByClientCIDR
type ServerAddressByClientCIDR struct {

	// The CIDR with which clients can match their IP to figure out the server address that they should use.
	// Required: true
	ClientCIDR *string `json:"clientCIDR"`

	// Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.
	// Required: true
	ServerAddress *string `json:"serverAddress"`
}
