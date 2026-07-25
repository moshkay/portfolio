package server

import "html/template"

// Portfolio is the top-level content model rendered into the site.
type Portfolio struct {
	Name       string
	Role       string
	Tagline    string
	About      []string
	Location   string
	Email      string
	Phone      string
	Socials    []Social
	Skills     []SkillGroup
	Projects   []Project
	Experience []Experience
	Year       int
}

// Social is an external profile link shown in the header/footer.
type Social struct {
	Label string
	URL   string
	Icon  template.HTML // inline SVG path data, rendered unescaped
}

// SkillGroup buckets related skills under a category heading.
type SkillGroup struct {
	Category string
	Items    []string
}

// Project is a single showcased piece of work.
type Project struct {
	Title       string
	Description string
	Tags        []string
	RepoURL     string
	LiveURL     string
	Featured    bool
}

// Experience is a role in the work-history timeline.
type Experience struct {
	Company string
	Role    string
	Period  string
	Summary string
}

// defaultPortfolio returns placeholder content. Replace these values with your
// own details, or wire this up to a database / CMS later.
func defaultPortfolio() Portfolio {
	return Portfolio{
		Name:     "Moshood Kolawole Saliu",
		Role:     "Software Engineer",
		Tagline:  "I build reliable backend systems and clean, fast web experiences.",
		Location: "Lagos, Nigeria",
		Email:    "kayzeebiz@gmail.com",
		Phone:    "+2349131943254",
		About: []string{
			"I'm a software engineer focused on backend systems — APIs, fintech wallet infrastructure, and identity verification services. I work primarily in Python and Go and enjoy turning fuzzy problems into simple, dependable systems.",
			"I've built multi-tenant SaaS platforms, real-time financial services, and KYC/identity verification tooling, with an emphasis on reliability, performance, and clean CI/CD.",
		},
		Socials: []Social{
			{Label: "GitHub", URL: "https://github.com/moshkay", Icon: iconGitHub},
			{Label: "LinkedIn", URL: "https://www.linkedin.com/in/saliu-moshood-a1b7421ab/", Icon: iconLinkedIn},
			{Label: "X", URL: "https://x.com/hemkayo", Icon: iconX},
		},
		Skills: []SkillGroup{
			{Category: "Languages", Items: []string{"Python", "Go", "JavaScript", "Java", "TypeScript", "SQL"}},
			{Category: "Backend", Items: []string{"Django", "FastAPI", "Flask", "Node.js", "REST APIs", "PostgreSQL", "Redis", "Kafka"}},
			{Category: "Data", Items: []string{"PostgreSQL", "MySQL"}},
			{Category: "Infra", Items: []string{"Docker", "GCP", "AWS", "Jenkins", "CI/CD", "Kubernetes", "Terraform", "Git"}},
		},
		Projects: []Project{
			{
				Title:       "Pila",
				Description: "A fintech platform providing virtual accounts, payment collections, transaction routing, and wallet infrastructure. Built the Django backend powering wallet management and real-time transaction processing.",
				Tags:        []string{"Django", "React", "Node.js"},
				RepoURL:     "",
				LiveURL:     "https://trypila.co",
				Featured:    true,
			},
			{
				Title:       "Dojah Verification Widget",
				Description: "A secure widget that lets businesses onboard customers through identity and document verification. Built and maintained Flask backend APIs supporting KYC and onboarding workflows.",
				Tags:        []string{"Flask", "Node.js", "EJS"},
				RepoURL:     "",
				LiveURL:     "https://identity.dojah.io",
				Featured:    true,
			},
			{
				Title:       "HRonWheels",
				Description: "A cloud-based, multi-tenant HR management platform for employee records, payroll, leave, attendance, and recruitment. Built the Django backend and cut API response times by ~20%.",
				Tags:        []string{"Django", "Python", "PostgreSQL", "DigitalOcean"},
				RepoURL:     "https://github.com/moshkay/hronwheels",
				LiveURL:     "https://www.hronwheels.ng",
				Featured:    false,
			},
			{
				Title:       "Portfolio Website",
				Description: "This site — a single self-contained Go binary that embeds its templates and assets. No external dependencies.",
				Tags:        []string{"Go", "net/http", "embed"},
				RepoURL:     "https://github.com/moshkay/portfolio",
				LiveURL:     "",
				Featured:    false,
			},
		},
		Experience: []Experience{
			{
				Company: "Dojah Inc.",
				Role:    "Software Engineer",
				Period:  "2020 — Present",
				Summary: "Deploy web applications and build backend services for identity verification and onboarding. Integrate multiple systems through APIs, drive code reviews and coding standards, and write unit tests to keep services reliable.",
			},
			{
				Company: "Elta Solutions",
				Role:    "Back-End Engineer",
				Period:  "2019 — 2020",
				Summary: "Built microservice-based architectures and secure RESTful APIs, and deployed applications on cloud platforms including AWS, Azure, and GCP.",
			},
			{
				Company: "Code Garage Africa",
				Role:    "Backend Developer",
				Period:  "2018 — 2019",
				Summary: "Developed robust APIs for HRonWheels, a multi-tenant HR platform, and improved back-end response times by ~20% on DigitalOcean infrastructure.",
			},
			{
				Company: "Longbridge Technologies",
				Role:    "Software Developer Intern",
				Period:  "2017",
				Summary: "Helped build a simplified deposit/withdrawal Finacle application for UBA Bank, cutting transaction completion time by ~50%.",
			},
		},
	}
}
