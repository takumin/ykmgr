package server

import (
	"context"
	"fmt"

	"github.com/go-piv/piv-go/piv"
	"github.com/takumin/ykmgr/pkg/yubikey/v1"
)

type server struct {
	yubikey.UnimplementedYubikeyServiceServer

	yk map[uint32]*piv.YubiKey
}

func NewServer() (*server, error) {
	cards, err := piv.Cards()
	if err != nil {
		return nil, err
	}

	yubikeys := make(map[uint32]*piv.YubiKey, len(cards))
	for _, v := range cards {
		handle, err := piv.Open(v)
		if err != nil {
			return nil, err
		}

		serial, err := handle.Serial()
		if err != nil {
			return nil, err
		}

		yubikeys[serial] = handle
	}

	return &server{
		yk: yubikeys,
	}, nil
}

func (s *server) Close() map[uint32]error {
	err := make(map[uint32]error, len(s.yk))
	for k, v := range s.yk {
		err[k] = v.Close()
	}
	return err
}

func (s *server) GetSerials(ctx context.Context, req *yubikey.GetSerialsRequest) (*yubikey.GetSerialsResponse, error) {
	keys := make([]uint32, 0, len(s.yk))
	for k := range s.yk {
		keys = append(keys, k)
	}
	return &yubikey.GetSerialsResponse{
		Serials: keys,
	}, nil
}

func (s *server) GetVersion(ctx context.Context, req *yubikey.GetVersionRequest) (*yubikey.GetVersionResponse, error) {
	var version yubikey.Version
	if _, ok := s.yk[req.GetSerial()]; ok {
		version.Major = uint32(s.yk[req.GetSerial()].Version().Major)
		version.Minor = uint32(s.yk[req.GetSerial()].Version().Minor)
		version.Patch = uint32(s.yk[req.GetSerial()].Version().Patch)
	}
	return &yubikey.GetVersionResponse{
		Version: &version,
	}, nil
}

func (s *server) GetRetries(ctx context.Context, req *yubikey.GetRetriesRequest) (*yubikey.GetRetriesResponse, error) {
	if _, ok := s.yk[req.GetSerial()]; !ok {
		return nil, fmt.Errorf("not found serial")
	}
	retries, err := s.yk[req.GetSerial()].Retries()
	if err != nil {
		return nil, err
	}
	return &yubikey.GetRetriesResponse{
		Retries: uint32(retries),
	}, nil
}

func (s *server) GetAttestationCertificate(ctx context.Context, req *yubikey.GetAttestationCertificateRequest) (*yubikey.GetAttestationCertificateResponse, error) {
	if _, ok := s.yk[req.GetSerial()]; !ok {
		return nil, fmt.Errorf("not found serial")
	}
	certificate, err := s.yk[req.GetSerial()].AttestationCertificate()
	if err != nil {
		return nil, err
	}
	return &yubikey.GetAttestationCertificateResponse{
		Certificate: certificate.Raw,
	}, nil
}
