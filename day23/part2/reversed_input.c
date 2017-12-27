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
		int f = 1;

		int d = 2;
		while (d != b) {
			int e = 2;
			while (e != b) {
				if (d * e == b) {
					f = 0;
				}
				e++;
			}
			d++;
		}

		if (f == 0) {
			h++;
		}

		if (b == c) {
			break;
		}

		b += 17;
	}

	printf("%d", h);
	return 0;
}
