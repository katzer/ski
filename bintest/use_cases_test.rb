# Write tests where the program should not fail
# this module is for testing the base/main use cases which are covered by this program
module UseCasesTest
  def test_pretty_print
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-p',
                                           '-d=true', 'app')
    check_error(output, error, 'pretty_print')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '|   0 | app       |', 'return was incorrect'
  end

  def test_script_execution
    output, error, status = Open3.capture3(PATH, BIN, '-s="test.sh"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_script_execution')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "bang\n", 'return was not correct'
  end

  def test_pretty_table_print
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="perlver_template"', '-p',
                                           '-d=true', 'app')
    check_error(output, error, 'test_pretty_tablePrint')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '| WILLYWONKA VERSION |', 'return was not right'
  end

  # TODO: Activate after fixing the fifa mock.
# def test_multiple_pretty_print
#   output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-p',
#                                          '-d=true', 'app', 'app', 'app')
#   check_error(output, error, 'pretty_print')
#   assert_true status.success?, 'Process did not exit cleanly'
#   assert_include output, '|   0 | app       |', 'return was incorrect'
#   assert_include output, '|   1 | app       |', 'return was incorrect'
#   assert_include output, '|   2 | app       |', 'return was incorrect'
# end

  def test_table_print
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="perlver_template"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_tablePrint')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "\n[\"willywonka_version\",", 'return was not right'
  end

  def test_empty_return
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo "',
                                           '-d=true', 'app')
    check_error(output, error, 'test_empty_return')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "\n", 'return was not empty'
  end
end
