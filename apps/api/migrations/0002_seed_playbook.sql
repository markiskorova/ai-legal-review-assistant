insert into playbook(org_id,name,description,active) values
(1,'Core SaaS Playbook','Baseline rules for SaaS MSA and NDA',true);

-- Examples
insert into rule(playbook_id,name,severity,pattern,guidance,llm_check) values
(1,'Liability Cap - 12mo fees','High','(?i)liabilit(y|ies).*(12|twelve).*(month)','Cap total liability to fees paid in the prior 12 months.',true),
(1,'Indemnity - IP Infringement','High','(?i)indemnif(y|ication).*(intellectual property|ip)','Require supplier to indemnify for IP infringement.',true),
(1,'Governing Law - Delaware','Medium','(?i)governing law','Prefer Delaware; flag others.',false),
(1,'Confidentiality','Medium','(?i)confidential','Ensure standard confidentiality obligations.',false);
