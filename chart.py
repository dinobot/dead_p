import datetime as dt
import json, pyc3, sys

dumpfile = sys.argv[1]
metric = sys.argv[2]

with open(dumpfile) as raw:
  data = json.load(raw)

x = ['x']
y = [metric]

for i in data:
  stamp = dt.datetime.utcfromtimestamp(float(i['timestamp'])).strftime("%Y-%m-%d")
  x.append(stamp)
  y.append(len(i['snapshot']))


pyc3.generate({

"data": {
        "x": 'x',
        "columns": [
            x,
            y 
        ]
    },
    "axis": {
        "x": {
            "type": 'timeseries',
            "tick": {
                "format": '%Y-%m-%d'
            }
        }
    }

})
