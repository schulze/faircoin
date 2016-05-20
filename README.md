# faircoin

Faircoin simulates a fair coin and should return 0 or 1 with equal probability.

In case a decision needs to be taken in an emergency, use

	$ seq 0 1 | shuf --random-source=/dev/urandom | head -n 1

instead.
