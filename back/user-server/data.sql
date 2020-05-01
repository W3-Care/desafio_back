--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2 (Debian 12.2-2.pgdg100+1)
-- Dumped by pg_dump version 12.1

-- Started on 2020-03-20 19:42:43 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2930 (class 0 OID 16387)
-- Dependencies: 203
-- Data for Name: marca; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.marca (id, nome) VALUES (1, 'Marca 1');
INSERT INTO public.marca (id, nome) VALUES (2, 'Marca 2');
INSERT INTO public.marca (id, nome) VALUES (3, 'Marca 3');


--
-- TOC entry 2932 (class 0 OID 16395)
-- Dependencies: 205
-- Data for Name: patrimonio; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.patrimonio (id, descricao, nome, marca_id) VALUES (1, 'Patrimonio 1', 'Ptm1', 1);


--
-- TOC entry 2934 (class 0 OID 16406)
-- Dependencies: 207
-- Data for Name: usuario; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.usuario (id, data_nascimento, email, nome, password) VALUES (1, '1984-10-14 21:00:00', 'bernardolobato@gmail.com', 'Bernardo Lobato', '$2a$10$g1PGTalqhfBObVw3t7GS0uS07yyE1hD2XHHfBwUwLS3b50LI4sZku');
INSERT INTO public.usuario (id, data_nascimento, email, nome, password) VALUES (2, '1984-10-14 21:00:00', 'admin@admin.com', 'Admin', '$2a$10$Q0cl7m.J3ijPKVi2wNhumePsvhRkSXdR0fd7DItqIKsNYDTEZciR.');


--
-- TOC entry 2940 (class 0 OID 0)
-- Dependencies: 202
-- Name: marca_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.marca_id_seq', 3, true);


--
-- TOC entry 2941 (class 0 OID 0)
-- Dependencies: 204
-- Name: patrimonio_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.patrimonio_id_seq', 1, true);


--
-- TOC entry 2942 (class 0 OID 0)
-- Dependencies: 206
-- Name: usuario_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.usuario_id_seq', 2, true);


-- Completed on 2020-03-20 19:42:43 UTC

--
-- PostgreSQL database dump complete
--

