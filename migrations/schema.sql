--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6
-- Dumped by pg_dump version 10.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: dostawcies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dostawcies (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.dostawcies OWNER TO postgres;

--
-- Name: fakturies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fakturies (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.fakturies OWNER TO postgres;

--
-- Name: kategories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kategories (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.kategories OWNER TO postgres;

--
-- Name: kliencis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kliencis (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.kliencis OWNER TO postgres;

--
-- Name: konta; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.konta (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.konta OWNER TO postgres;

--
-- Name: platnoscis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.platnoscis (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.platnoscis OWNER TO postgres;

--
-- Name: pozycjes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pozycjes (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.pozycjes OWNER TO postgres;

--
-- Name: produkties; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.produkties (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.produkties OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: sprzedawcies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sprzedawcies (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.sprzedawcies OWNER TO postgres;

--
-- Name: zamowienia; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.zamowienia (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.zamowienia OWNER TO postgres;

--
-- Name: dostawcies dostawcies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dostawcies
    ADD CONSTRAINT dostawcies_pkey PRIMARY KEY (id);


--
-- Name: fakturies fakturies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fakturies
    ADD CONSTRAINT fakturies_pkey PRIMARY KEY (id);


--
-- Name: kategories kategories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kategories
    ADD CONSTRAINT kategories_pkey PRIMARY KEY (id);


--
-- Name: kliencis kliencis_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kliencis
    ADD CONSTRAINT kliencis_pkey PRIMARY KEY (id);


--
-- Name: konta konta_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.konta
    ADD CONSTRAINT konta_pkey PRIMARY KEY (id);


--
-- Name: platnoscis platnoscis_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.platnoscis
    ADD CONSTRAINT platnoscis_pkey PRIMARY KEY (id);


--
-- Name: pozycjes pozycjes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pozycjes
    ADD CONSTRAINT pozycjes_pkey PRIMARY KEY (id);


--
-- Name: produkties produkties_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.produkties
    ADD CONSTRAINT produkties_pkey PRIMARY KEY (id);


--
-- Name: sprzedawcies sprzedawcies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sprzedawcies
    ADD CONSTRAINT sprzedawcies_pkey PRIMARY KEY (id);


--
-- Name: zamowienia zamowienia_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zamowienia
    ADD CONSTRAINT zamowienia_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

