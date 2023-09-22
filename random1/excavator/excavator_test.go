package excavator_test

import (
	"log/slog"
	"testing"

	"github.com/BarunW/random/excavator"
)

func TestMakeRequest(t *testing.T) {
    ne := excavator.NewExcavator(&slog.Logger{})    
    
    err := ne.StartExcavating()
    if err != nil{
        t.Fatal("Error", err)
    }
}
