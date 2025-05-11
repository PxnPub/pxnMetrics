package main;

import(
	Fmt    "fmt"
	Net    "net"
	Time   "time"
	Atomic "sync/atomic"
	"strconv"
);



const ThreadCount = 1;



func main() {
	print("\nStarting..\n\n");
	var count int64 = 0;

	tick_interval, _ := Time.ParseDuration("1s");
	ticker := Time.NewTicker(tick_interval);
	defer ticker.Stop();
	go func() {
		for {
			select {
			case <-ticker.C:
				cnt := count;
				Atomic.AddInt64(&count, -cnt);
				Fmt.Printf(" %s per sec\n", Format(cnt));
			}
		}
	}();

	address := "127.0.0.1:9001";

	for i:=0; i<ThreadCount; i++ {
		go func() {
			addr, err := Net.ResolveUDPAddr("udp", address)
			if err != nil { panic(err); }
			conn, err := Net.DialUDP("udp", nil, addr);
			if err != nil { panic(err); }
			for {
				Atomic.AddInt64(&count, 1);
				msg := []byte("{}");
				_, err := conn.Write(msg);
				if err != nil { panic(err); }
			}
		}();
	}

	sleep, _ := Time.ParseDuration("5s");
	for {
		Time.Sleep(sleep);
	}
	print("<end>\n\n");
}



func Format(n int64) string {
    in := strconv.FormatInt(n, 10)
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
