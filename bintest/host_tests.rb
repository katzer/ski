# Tests for the target systems: supported and unsupported types, invalid planets etc.
module HostTests
  def test_web
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
                                           '-d=true', 'web')
    check_no_error(output, error, 'test_web')
    assert_true status.success?, 'Process did exit cleanly'
    assert_include error, 'Usage of ski with web servers is not implemented'
  end

  def test_server
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_server')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '123'
  end

  def test_not_authorized_host
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
                                          '-d=true', 'unauthorized')
    check_no_error(output, error, 'test_not_authorized_host')
    assert_true status.success?, 'Process did exit cleanly'
    # NOTE: error output depends on the fifa implementation/version so don't check it.
  end

  def test_nonexistent_planet
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-d=true',
                                           'offline')
    check_no_error(output, error, 'nonexistent_planet')
    assert_true status.success?, 'Process did exit cleanly'
  end
end
