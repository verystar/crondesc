package crondesc

func da_TestCases() []localeTestCase {
	return []localeTestCase{
		// TODO: Need help
		{inExpr: "* * * * *", outErr: nil, outDesc: "Hvert minut"},
	}
}
