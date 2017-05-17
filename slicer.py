import json, sys

tssnap = json.load(sys.stdin)

for t in tssnap:
  with open('slices/'+t['timestamp'], 'w') as file:
    file.write(json.dumps(t['snapshot']))


