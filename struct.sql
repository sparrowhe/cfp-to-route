/*
 Navicat Premium Data Transfer

 Source Server         : Navdata
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 08/07/2022 14:18:31
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for airports
-- ----------------------------
DROP TABLE IF EXISTS "airports";
CREATE TABLE "airports" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "icao" TEXT,
  "latitude" text,
  "longitude" TEXT
);

-- ----------------------------
-- Table structure for airways
-- ----------------------------
DROP TABLE IF EXISTS "airways";
CREATE TABLE "airways" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "leg_id" INTEGER,
  "name" TEXT,
  "point" TEXT,
  "latitude" TEXT,
  "longitude" TEXT
);

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "sqlite_sequence";
CREATE TABLE "sqlite_sequence" (
  "name",
  "seq"
);

-- ----------------------------
-- Table structure for waypoints
-- ----------------------------
DROP TABLE IF EXISTS "waypoints";
CREATE TABLE "waypoints" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "latitude" TEXT,
  "longitude" TEXT
);

-- ----------------------------
-- Indexes structure for table airways
-- ----------------------------
CREATE INDEX "airways_legid_name_index"
ON "airways" (
  "leg_id" ASC,
  "name" ASC
);

-- ----------------------------
-- Auto increment value for waypoints
-- ----------------------------

PRAGMA foreign_keys = true;
