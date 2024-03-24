package parser_utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func check(t *testing.T, stmt string, outQueryExpected string, outColumnsExpected []string) {
	query, columns, err := ExtractQueryAndColumns(stmt)
	require.NoError(t, err)
	require.Equal(t, outQueryExpected, query)
	require.Equal(t, outColumnsExpected, columns)
}

func TestExtractQueryAndColumns(t *testing.T) {
	check(
		t,
		"INSERT INTO `my_schema`.`my_table` (`__primary_key`,`ColumnNameWithParentheses(something)`) VALUES (1,1);",
		"INSERT INTO `my_schema` . `my_table` ( `__primary_key` , `ColumnNameWithParentheses(something)` ) VALUES ",
		[]string{"__primary_key", "ColumnNameWithParentheses(something)"},
	)
}
