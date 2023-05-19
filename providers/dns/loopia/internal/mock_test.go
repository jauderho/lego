package internal

const (
	exampleDomain    = "example.com"
	exampleSubDomain = "_acme-challenge"
	exampleRdata     = "LHDhK3oGRvkiefQnx7OOczTY5Tic_xZ6HcMOc_gmtoM"
)

// Testdata based on real traffic between a xml-rpc client and the api.
const responseOk = `<?xml version="1.0" encoding="UTF-8"?>
		<methodResponse>
		  <params>
			<param>
			  <value>
				<string>
				  OK
				</string>
			  </value>
			</param>
		  </params>
		</methodResponse>`

const responseAuthError = `<?xml version="1.0" encoding="UTF-8"?>
		<methodResponse>
		  <params>
			<param>
			  <value>
				<string>
				  AUTH_ERROR
				</string>
			  </value>
			</param>
		  </params>
		</methodResponse>`

const responseUnknownError = `<?xml version="1.0" encoding="UTF-8"?>
		<methodResponse>
		  <params>
			<param>
			  <value>
				<string>
				  UNKNOWN_ERROR
				</string>
			  </value>
			</param>
		  </params>
		</methodResponse>`

const responseRPCError = `<?xml version="1.0" encoding="UTF-8"?>
<methodResponse>
  <fault>
    <value>
      <struct>
        <member>
          <name>
            faultCode
          </name>
          <value>
            <int>
              201
            </int>
          </value>
        </member>
        <member>
          <name>
            faultString
          </name>
          <value>
            <string>
              Method signature error: 42
            </string>
          </value>
        </member>
      </struct>
    </value>
  </fault>
</methodResponse>`

const addZoneRecordGoodAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>addZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <struct>
          <member>
            <name>type</name>
            <value>
              <string>TXT</string>
            </value>
          </member>
          <member>
            <name>ttl</name>
            <value>
              <int>123</int>
            </value>
          </member>
          <member>
            <name>priority</name>
            <value>
              <int>0</int>
            </value>
          </member>
          <member>
            <name>rdata</name>
            <value>
              <string>TXTrecord</string>
            </value>
          </member>
          <member>
            <name>record_id</name>
            <value>
              <int>0</int>
            </value>
          </member>
        </struct>
      </value>
    </param>
  </params>
</methodCall>`

const addZoneRecordBadAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>addZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>badpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <struct>
          <member>
            <name>type</name>
            <value>
              <string>TXT</string>
            </value>
          </member>
          <member>
            <name>ttl</name>
            <value>
              <int>123</int>
            </value>
          </member>
          <member>
            <name>priority</name>
            <value>
              <int>0</int>
            </value>
          </member>
          <member>
            <name>rdata</name>
            <value>
              <string>TXTrecord</string>
            </value>
          </member>
          <member>
            <name>record_id</name>
            <value>
              <int>0</int>
            </value>
          </member>
        </struct>
      </value>
    </param>
  </params>
</methodCall>`

const addZoneRecordNonValidDomain = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>addZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>badexample.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <struct>
          <member>
            <name>type</name>
            <value>
              <string>TXT</string>
            </value>
          </member>
          <member>
            <name>ttl</name>
            <value>
              <int>123</int>
            </value>
          </member>
          <member>
            <name>priority</name>
            <value>
              <int>0</int>
            </value>
          </member>
          <member>
            <name>rdata</name>
            <value>
              <string>TXTrecord</string>
            </value>
          </member>
          <member>
            <name>record_id</name>
            <value>
              <int>0</int>
            </value>
          </member>
        </struct>
      </value>
    </param>
  </params>
</methodCall>`

const addZoneRecordEmptyResponse = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>addZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>empty.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <struct>
          <member>
            <name>type</name>
            <value>
              <string>TXT</string>
            </value>
          </member>
          <member>
            <name>ttl</name>
            <value>
              <int>123</int>
            </value>
          </member>
          <member>
            <name>priority</name>
            <value>
              <int>0</int>
            </value>
          </member>
          <member>
            <name>rdata</name>
            <value>
              <string>TXTrecord</string>
            </value>
          </member>
          <member>
            <name>record_id</name>
            <value>
              <int>0</int>
            </value>
          </member>
        </struct>
      </value>
    </param>
  </params>
</methodCall>`

const getZoneRecords = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>getZoneRecords</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
  </params>
</methodCall>`

const getZoneRecordsResponse = `<?xml version="1.0" encoding="UTF-8"?>
<methodResponse>
  <params>
    <param>
      <value>
        <array>
          <data>
            <value>
              <struct>
                <member>
                  <name>
                    rdata
                  </name>
                  <value>
                    <string>
                    LHDhK3oGRvkiefQnx7OOczTY5Tic_xZ6HcMOc_gmtoM
                    </string>
                  </value>
                </member>
                <member>
                  <name>
                    record_id
                  </name>
                  <value>
                    <int>
                      12345678
                    </int>
                  </value>
                </member>
                <member>
                  <name>
                    priority
                  </name>
                  <value>
                    <int>
                      0
                    </int>
                  </value>
                </member>
                <member>
                  <name>
                    ttl
                  </name>
                  <value>
                    <int>
                      300
                    </int>
                  </value>
                </member>
                <member>
                  <name>
                    type
                  </name>
                  <value>
                    <string>
                      TXT
                    </string>
                  </value>
                </member>
              </struct>
            </value>
          </data>
        </array>
      </value>
    </param>
  </params>
</methodResponse>`

const removeRecordGoodAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <int>12345678</int>
      </value>
    </param>
  </params>
</methodCall>`

const removeRecordBadAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>badpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <int>12345678</int>
      </value>
    </param>
  </params>
</methodCall>`

const removeRecordNonValidDomain = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>badexample.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <int>12345678</int>
      </value>
    </param>
  </params>
</methodCall>`

const removeRecordEmptyResponse = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeZoneRecord</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>empty.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
    <param>
      <value>
        <int>12345678</int>
      </value>
    </param>
  </params>
</methodCall>`

const removeSubdomainGoodAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeSubdomain</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
  </params>
</methodCall>`

const removeSubdomainBadAuth = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeSubdomain</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>badpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>example.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
  </params>
</methodCall>`

const removeSubdomainNonValidDomain = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeSubdomain</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>badexample.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
  </params>
</methodCall>`

const removeSubdomainEmptyResponse = `<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>removeSubdomain</methodName>
  <params>
    <param>
      <value>
        <string>apiuser</string>
      </value>
    </param>
    <param>
      <value>
        <string>goodpassword</string>
      </value>
    </param>
    <param>
      <value>
        <string>empty.com</string>
      </value>
    </param>
    <param>
      <value>
        <string>_acme-challenge</string>
      </value>
    </param>
  </params>
</methodCall>`
