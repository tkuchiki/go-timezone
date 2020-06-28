import json
import sys
from jinja2 import Template

def var_tzabbrinfos():
    data = {}
    data2 = {}
    abbrs_json = sys.argv[1]
    military_abbrs_json = sys.argv[2]
    with open(abbrs_json) as f:
        data = json.load(f)

    with open(military_abbrs_json) as f:
        data2 = json.load(f)

    tpl_text = '''var tzAbbrInfos = map[string][]*TzAbbreviationInfo{
    {%- for abbr in tzinfos %}
    "{{ abbr }}": []*TzAbbreviationInfo{
        {%- for tz in tzinfos[abbr] %}
        {
            countryCode: "{{ tz["country_code"] }}",
            isDST: {{ tz["is_dst"] | lower }},
            name: "{{ tz["name"] }}",
            offset: {{ tz["offset"] }},
            offsetHHMM: "{{ tz["offset_hhmm"] }}",
        },
        {%- endfor %}
    },
    {%- endfor %}
    // military timezones
    {%- for abbr in military_tzinfos %}
    "{{ abbr }}": []*TzAbbreviationInfo{
        {
            name: "{{ military_tzinfos[abbr]["name"] }}",
            offset: {{ military_tzinfos[abbr]["offset"] }},
            offsetHHMM: "{{ military_tzinfos[abbr]["offset_hhmm"] }}",
        },
    },
    {%- endfor %}
}'''

    tpl = Template(tpl_text)

    print(tpl.render({"tzinfos": data, "military_tzinfos": data2}))

if __name__ == "__main__":
    var_tzabbrinfos()