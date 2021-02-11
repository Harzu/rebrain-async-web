01:
	@go run ./cmd/01_syncmode

# Goroutines
02:
	@go run ./cmd/goroutine/02_nowait

03:
	@go run ./cmd/goroutine/03_sleepwait

04:
	@go run ./cmd/goroutine/04_runjobs

05:
	@go run ./cmd/goroutine/05_funcparams

06:
	@go run ./cmd/goroutine/06_mainalsogoroutine

# Channels
07:
	@go run ./cmd/channels/07_sendreceive

08:
	@go run ./cmd/channels/08_senddl

09:
	@go run ./cmd/channels/09_receivedl

10:
	@go run ./cmd/channels/10_concurrentsend

# Buffer channels
11:
	@go run ./cmd/channels/bufferchannel/11_send

12:
	@go run ./cmd/channels/bufferchannel/12_sendbufoverflow

13:
	@go run ./cmd/channels/bufferchannel/13_receive

14:
	@go run ./cmd/channels/bufferchannel/14_concurrentsend

# Select
15:
	@go run ./cmd/channels/select/15_pingpong

16:
	@go run ./cmd/channels/select/16_withoutvalue

17:
	@go run ./cmd/channels/select/17_default

18:
	@go run ./cmd/channels/close/18_sendtoclosechannel

19:
	@go run ./cmd/channels/close/19_receivefromclosechannel

# Synchronization
20:
	@go run ./cmd/syncronization/20_race

21:
	@go run ./cmd/syncronization/21_map

22:
	@go run ./cmd/syncronization/22_mutex

23:
	@go run ./cmd/syncronization/23_rwmutex

24:
	@go run ./cmd/syncronization/24_waitgroup

# Context
25:
	@go run ./cmd/context/25_cancel

26:
	@go run ./cmd/context/26_timeout

27:
	@go run ./cmd/context/27_value