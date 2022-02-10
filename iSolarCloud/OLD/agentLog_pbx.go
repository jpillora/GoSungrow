package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountAgentLogs(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query, err := p.Areas.AgentLog.Count(domain)
		if err != nil {
			p.Error = err
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("AgentLogs: %v\n", query.Response.Total)
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
	}

	return p.Error
}

func (p *SunGro) ListAgentLogs(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query, err := p.Areas.AgentLog.List(domain)
		if err != nil {
			p.Error = err
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("AgentLogs:\n%v\n", query.Response.String())
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
		if p.OutputType == TypeGoogle {
			p.OutputArray = query.Response.ToArray()
			// data := query.Response.ToArray()
			// p.Error = p.UpdateGoogleSheet("device", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadAgentLogs(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query, err := p.Areas.AgentLog.Read(domain)
		if err != nil {
			p.Error = err
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("AgentLogs:\n%v\n", query.Response.String())
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
		if p.OutputType == TypeGoogle {
			p.OutputArray = query.Response.ToArray()
			// data := query.Response.ToArray()
			// p.Error = p.UpdateGoogleSheet("device", data)
			break
		}
	}

	return p.Error
}