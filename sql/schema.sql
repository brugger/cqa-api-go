
CREATE TABLE instrument (
  id    UUID NOT NULL DEFAULT  uuid_generate_v4 () PRIMARY KEY,
  name  VARCHAR(80) NOT NULL
);


CREATE TABLE runfolder (
  id            UUID NOT NULL DEFAULT  uuid_generate_v4 () PRIMARY KEY,
  name          VARCHAR(80) NOT NULL UNIQUE,
  run_date      timestamptz,

  samples     	       INT,
  total_reads	       INT,
  mapped_reads	       INT,
  duplicate_reads      INT,
  mean_isize	       float
);



CREATE TABLE sample (
  id            UUID NOT NULL DEFAULT  uuid_generate_v4 () PRIMARY KEY,

  name          VARCHAR(80) NOT NULL,
  runfolder_id  UUID references runfolder(id),

  UNIQUE  ("name", "runfolder_id")
);

CREATE INDEX sampe_name_idx ON sample (
  name
);




CREATE TABLE fq_stats (
  id  UUID NOT NULL DEFAULT  uuid_generate_v4 () PRIMARY KEY,

  sample_id UUID NOT NULL references  sample( id ),

  total_reads	       INT,
  gc_perc              float,
  seq_quality_mean     float,
  seq_quality_95p      float,
  seq_length_mean      float,
  seq_length_95p       float
);


