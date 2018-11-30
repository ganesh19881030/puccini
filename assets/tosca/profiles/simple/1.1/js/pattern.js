
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.2

function validate(v, re) {
	if (arguments.length !== 2)
		throw 'must have 1 argument';
	if (v.$string !== undefined)
		v = v.$string;
	return new RegExp(re).test(v);
}
