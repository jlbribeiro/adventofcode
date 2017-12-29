#include <stdio.h>

/* This is the program obtained by reverse-engineering the input file.
 * It has no obvious optimizations _on purpose_; it aims to be an easy-to-read
 * version of the input file.
 */
int main(int argc, char **argv) {
	int b = 99 * 100 + 100000;
	int c = b + 17000;
	int h = 0;

	while (1) {
		int d = 2;
		while (d != b) {
			if (b % d == 0) {
				h++;
				break;
			}
			d++;
		}

		if (b == c) {
			break;
		}

		b += 17;
	}

	printf("%d\n", h);
	return 0;
}
