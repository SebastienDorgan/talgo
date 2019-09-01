#!/usr/bin/python

from string import Template


def get_types(base, sizes=["", "8", "16", "32", "64"]):
    return ["{}{}".format(base, s) for s in sizes]

types = []
types.extend(get_types("uint"))
types.extend(get_types("int"))
types.extend(get_types("float", ["32","64"]))
print(types)
tpl = Template("""
package talgo


//Greater check if s[i] is greater or greater or equal to value
func ${Type}Greater(s []${type}, value ${type}, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func  ${Type}Lesser(s []${type}, value ${type}, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func ${Type}Equal(s []${type}, value ${type}) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
"""
)

def generate(type):
    s =tpl.substitute(type=type, Type=type.title())
    with open("{}predicates.go".format(type),"w") as f:
        f.write(s)

for type in types:
    generate(type)


