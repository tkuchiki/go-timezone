import json
import sys
from jinja2 import Template

def var_timezones():
    data = {}
    abbr_timezones_json = sys.argv[1]
    with open(abbr_timezones_json) as f:
        data = json.load(f)

    tpl_text = '''var timezones = map[string][]string{
    {%- for abbr in timezones %}
    "{{ abbr }}": []string{
        {%- for tz in timezones[abbr] %}
        "{{ tz }}",
        {%- endfor %}
    },
    {%- endfor %}
}'''

    tpl = Template(tpl_text)

    print(tpl.render({"timezones": data}))

if __name__ == "__main__":
    var_timezones()