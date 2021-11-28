#/usr/bin/env bash

_wlog_completions() {
	_get_comp_words_by_ref -n : cur

	if [ "${COMP_WORDS[1]}" == "start" ]; then
		COMPREPLY=($(wlog find "${cur}"))
		__ltrim_colon_completions "${cur}"
		return
	fi

	if [ "${COMP_WORDS[1]}" == "break" ]; then
		COMPREPLY=($(compgen -W "start end" "${cur}"))
	elif [ "${COMP_WORDS[1]}" == "report" ]; then
		COMPREPLY=($(compgen -W "day week" "${cur}"))
	else
		COMPREPLY=($(compgen -W "break end in non out status timesheet report start" "${cur}"))
	fi

}

complete -F _wlog_completions wlog
