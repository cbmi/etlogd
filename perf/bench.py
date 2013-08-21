#!/usr/bin/env python

import re
import os
import sys
import time
from cStringIO import StringIO
import subprocess

ab_cmd = 'ab -T application/json -k -p payload.json -c {0} -n {1} http://0.0.0.0:4000/'

errors_re = re.compile(r'Failed requests:\s+(\d+)')
rate_re = re.compile(r'Requests per second:\s+(\d+)')
sock_re = re.compile(r'apr_socket_recv:')

output_template = '{requests}:{concurrency}\t{rate}/s\t{errors}'

REQUEST_STEPS = (1000, 5000)
CONCONCURRENCY_STEPS = (1, 5, 10, 25, 50, 100, 200, 500, 1000)

FULL_MATRIX = [(r, c) for r in REQUEST_STEPS for c in CONCONCURRENCY_STEPS]

def ab_bench(r, c):
    fname = 'tmp-results.txt'
    devnull = open(os.devnull, 'w')

    with open(fname, 'w') as buff:
        code = subprocess.call(ab_cmd.format(c, r).split(), stdout=buff, stderr=devnull)

    kwargs = {
        'requests': r,
        'concurrency': c,
    }

    with open(fname) as buff:
        output = buff.read()

        if code > 0:
            match = sock_re.search(output)
            if match:
                error = match.groups()[0]
            else:
                error = '(no message supplied)'
            print '{requests}:{concurrency} - error! {error}'.format(error=error, **kwargs)
        else:
            kwargs['errors'] = errors_re.search(output).groups()[0]
            kwargs['rate'] = rate_re.search(output).groups()[0]
            print output_template.format(**kwargs)

    os.remove(fname)

if __name__ == '__main__':
    import sys

    steps = [a.split('/') for a in sys.argv[1:]]

    if not steps:
        steps = FULL_MATRIX

    try:
        for step in steps:
            ab_bench(*step)
            time.sleep(1)
    except KeyboardInterrupt:
        sys.stdout.write('\rStopped.\n')
