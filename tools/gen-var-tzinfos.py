import json
import sys
from jinja2 import Template

def var_tzinfos():
    data = {}
    timezones_json = sys.argv[1]
    with open(timezones_json) as f:
        data = json.load(f)

    tpl_text = '''var tzInfos = map[string]*TzInfo{
    {%- for tz in tzinfos %}
    "{{ tz }}": &TzInfo{
        longGeneric: "{{ tzinfos[tz]["long"]["generic"] }}",
        longStandard: "{{ tzinfos[tz]["long"]["standard"] }}",
        longDaylight: "{{ tzinfos[tz]["long"]["daylight"] }}",
        shortGeneric: "{{ tzinfos[tz]["short"]["generic"] }}",
        shortStandard: "{{ tzinfos[tz]["short"]["standard"] }}",
        shortDaylight: "{{ tzinfos[tz]["short"]["daylight"] }}",
        standardOffset: {{ tzinfos[tz]["standard_offset"] }},
        daylightOffset: {{ tzinfos[tz]["daylight_offset"] }},
        standardOffsetHHMM: "{{ tzinfos[tz]["standard_offset_hhmm"] }}",
        daylightOffsetHHMM: "{{ tzinfos[tz]["daylight_offset_hhmm"] }}",
        countryCode: "{{ tzinfos[tz]["country_code"] }}",
        isDeprecated: {{ tzinfos[tz]["is_deprecated"] | lower }},
        linkTo: "{{ tzinfos[tz]["link_to"] }}",
        lastDST: {{ tzinfos[tz]["last_dst"] }},
    },
    {%- endfor %}
}'''

    tpl = Template(tpl_text)

    print(tpl.render({"tzinfos": data}))

if __name__ == "__main__":
    var_tzinfos()