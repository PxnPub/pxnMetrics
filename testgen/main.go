package main;

import(
	OS      "os"
	Fmt     "fmt"
	Net     "net"
	Time    "time"
	Rand    "math/rand"
	StrConv "strconv"
	Atomic  "sync/atomic"
	JSON    "encoding/json"
	API     "github.com/PxnPub/pxnMetrics/api/submitapi"
);



//const ThreadCount = 200;
const ThreadCount = 100;
//const Sleep = "50ms";



func main() {
	print("\nStarting..\n\n");
	var count Atomic.Uint64;
	var last  Atomic.Uint64;

	tick_interval, _ := Time.ParseDuration("1s");
	ticker := Time.NewTicker(tick_interval);
	defer ticker.Stop();
	go func() {
		for {
			select {
			case <-ticker.C:
				cnt := count.Load();
				Fmt.Printf(" %s per sec\n", Format(int64(cnt-last.Load())));
				last.Store(cnt);
			}
		}
	}();

	address := "127.0.0.1:9001";
//sleep, _ := Time.ParseDuration(Sleep);
	for i:=0; i<ThreadCount; i++ {
		go func() {
			addr, err := Net.ResolveUDPAddr("udp", address)
			if err != nil { panic(err); }
			conn, err := Net.DialUDP("udp", nil, addr);
			if err != nil { panic(err); }
			for {
//Time.Sleep(sleep);
sleep := Time.Duration(Rand.Intn(1000)) * Time.Millisecond;
//Fmt.Printf("Sleep: %s\n", sleep);
Time.Sleep(sleep);
				count.Add(1);
				// build submit packet
				timestamp := Time.Now().UnixMilli();
				uid := uint64(Rand.Intn(7999999999999999999));
				num_players := int16(Rand.Intn(123));
				var platform string;
				switch Rand.Intn(7) {
					case 0, 1, 2, 3: platform = "PaperMC"; break;
					case 4:          platform = "Folia";   break;
					case 5, 6, 7:    platform = "Fabric";  break;
				}
				// packet
				json, err := JSON.Marshal(API.Submit{
					Timestamp:  timestamp,
					ServerUID:  uid,
					Platform:   platform,
					NumPlayers: num_players,
				});
				if err != nil { panic(err); }
				// send
				_, err = conn.Write([]byte(json));
				if err != nil { panic(err); }
			}
		}();
	}

	slp, _ := Time.ParseDuration("5.1s");
	for {
		Time.Sleep(slp);
	}
	print("<end>\n\n");
	OS.Exit(0);
}



func Format(n int64) string {
    in := StrConv.FormatInt(n, 10)
    numOfDigits := len(in)
    if n < 0 {
        numOfDigits-- // First character is the - sign (not a digit)
    }
    numOfCommas := (numOfDigits - 1) / 3

    out := make([]byte, len(in)+numOfCommas)
    if n < 0 {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}
