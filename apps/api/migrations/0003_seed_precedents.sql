-- 0003_seed_precedents.sql
-- Seed sample precedent clauses for testing search & review

INSERT INTO precedent (org_id, title, text, tags)
VALUES
  (
    1,
    'NDA Confidentiality',
    'Each party shall maintain in strict confidence all confidential information disclosed by the other party and shall not disclose such information to any third party without prior written consent.',
    'confidentiality,nda,us'
  ),
  (
    1,
    'Limitation of Liability',
    'Except in cases of gross negligence or willful misconduct, neither party shall be liable for any indirect, incidental, or consequential damages arising out of this agreement.',
    'liability,damages,us'
  ),
  (
    1,
    'Governing Law (California)',
    'This Agreement shall be governed by and construed in accordance with the laws of the State of California, without regard to its conflict of law provisions.',
    'governing law,california,us'
  ),
  (
    1,
    'Force Majeure',
    'Neither party shall be liable for any failure or delay in performance due to causes beyond its reasonable control, including acts of God, natural disasters, or government actions.',
    'force majeure,global'
  );
