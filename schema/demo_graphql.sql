/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : PostgreSQL
 Source Server Version : 100400
 Source Host           : localhost
 Source Database       : demo_graphql
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100400
 File Encoding         : utf-8
*/

-- ----------------------------
--  Table structure for stores
-- ----------------------------
DROP TABLE IF EXISTS "public"."stores";
CREATE TABLE "public"."stores" (
	"id" int4 NOT NULL,
	"name" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."stores" OWNER TO "postgres";

-- ----------------------------
--  Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "public"."products";
CREATE TABLE "public"."products" (
	"id" int4 NOT NULL,
	"store_id" int4 NOT NULL,
	"name" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."products" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table stores
-- ----------------------------
ALTER TABLE "public"."stores" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

