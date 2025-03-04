resource "citrixadc_aaauser_vpntrafficpolicy_binding" "tf_aaauser_vpntrafficpolicy_binding" {
  username = "user1"
  policy    = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
  priority  = 100
}

resource "citrixadc_vpntrafficaction" "foo" {
  fta        = "ON"
  hdx        = "ON"
  name       = "Testingaction"
  qual       = "tcp"
  sso        = "ON"
}
resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
  name   = "tf_vpntrafficpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpntrafficaction.foo.name
}