# This module fis or tests where the program becomes the correct number of arguments but
# where the arguments are illegal.
module IllegalArgumentsTest

  def test_bad_command
    output, error, status = Open3.capture3(PATH, BIN, '-c="yabeda baba"',
                                           '-d=true', 'app')
    check_no_error(output, error, 'bad_command')
    assert_include output, 'Process exited with status 127', 'return incorrect'
    # assert_false status.success?, 'Process did exit cleanly' # TODO check why status is ok
  end

  def test_malformed_template
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="useless_template"', '-d=true',
                                           '-p', 'app')
    check_error(output, error, 'malformed_template')
    # assert_true status.success?, 'Process did exit cleanly' # TODO check why status is ok
    # assert_include error, 'exit status 2', 'wrong error'
  end

  def test_no_template
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="no_template"', '-d=true',
                                           '-p', 'app')
    check_no_error(output, error, 'no_template')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'The provided template does not exist', 'wrong error'
  end

  def test_bad_script
    output, error, status = Open3.capture3(PATH, BIN, '-s="badscript.sh"',
                                           'app')
    check_error(output, error, 'bad_script')
    assert_include output, 'Process exited with status 127', 'return incorrect'
    # assert_false status.success?, 'Process did exit cleanly' # TODO check why status is ok
  end

  def test_no_such_script
    output, error, status = Open3.capture3(PATH, BIN, '-s="nonExistent.sh"',
                                           '-d=true', 'app')
    check_no_error(output, error, 'no_such_script')
    assert_include output, 'no such file or directory', 'error was not correct'
    # assert_false status.success?, 'Process did exit cleanly' # TODO check why status is ok
  end
end
