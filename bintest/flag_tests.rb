# Test the program different combinations of the flags and different values
module FlagTests
  def test_not_enough_args
    output, error, status = Open3.capture3(PATH, BIN, '-p', '-d=true', 'app')
    check_error(output, error, 'not_enough_args')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage:', 'return was not correct'
  end

  def test_wrong_flag_order
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', 'app',
                                           '-d=true', '-p')
    check_no_error(output, error, 'wrong_flag_order')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'Unknown target', 'error was not correct'
  end

  def test_malformed_flag
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-zz',
                                           '-d=true', 'app')
    check_no_error(output, error, 'malformed_flag')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'Usage of', 'return was not correct'
  end

  def test_version
    output, error, status = Open3.capture3(PATH, BIN, '-v')
    check_error(output, error, 'test_version')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '0.9'
  end
end
