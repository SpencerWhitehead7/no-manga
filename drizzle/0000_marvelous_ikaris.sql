CREATE TABLE `chapter` (
	`manga_id` integer NOT NULL,
	`chapter_num` real NOT NULL,
	`name` text,
	`page_count` integer NOT NULL,
	`updated_at` integer DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`chapter_num`, `manga_id`),
	FOREIGN KEY (`manga_id`) REFERENCES `manga`(`id`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `demo` (
	`name` text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE `genre` (
	`name` text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE `job` (
	`name` text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE `magazine` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text NOT NULL,
	`other_names` text NOT NULL,
	`description` text NOT NULL,
	`demo` text NOT NULL,
	FOREIGN KEY (`demo`) REFERENCES `demo`(`name`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `magazine_manga` (
	`magazine_id` integer NOT NULL,
	`manga_id` integer NOT NULL,
	FOREIGN KEY (`magazine_id`) REFERENCES `magazine`(`id`) ON UPDATE no action ON DELETE no action,
	FOREIGN KEY (`manga_id`) REFERENCES `manga`(`id`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `manga` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text NOT NULL,
	`other_names` text NOT NULL,
	`description` text NOT NULL,
	`demo` text NOT NULL,
	`start_date` integer NOT NULL,
	`end_date` integer,
	FOREIGN KEY (`demo`) REFERENCES `demo`(`name`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `manga_genre` (
	`manga_id` integer NOT NULL,
	`genre` text NOT NULL,
	FOREIGN KEY (`manga_id`) REFERENCES `manga`(`id`) ON UPDATE no action ON DELETE no action,
	FOREIGN KEY (`genre`) REFERENCES `genre`(`name`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `manga_mangaka_job` (
	`manga_id` integer NOT NULL,
	`mangaka_id` integer NOT NULL,
	`job` text NOT NULL,
	FOREIGN KEY (`manga_id`) REFERENCES `manga`(`id`) ON UPDATE no action ON DELETE no action,
	FOREIGN KEY (`mangaka_id`) REFERENCES `mangaka`(`id`) ON UPDATE no action ON DELETE no action,
	FOREIGN KEY (`job`) REFERENCES `job`(`name`) ON UPDATE no action ON DELETE no action
);
--> statement-breakpoint
CREATE TABLE `mangaka` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text NOT NULL,
	`other_names` text NOT NULL,
	`description` text NOT NULL
);
--> statement-breakpoint
CREATE INDEX `idx_chapter_manga_id` ON `chapter` (`manga_id`);--> statement-breakpoint
CREATE INDEX `idx_chapter_updated_at` ON `chapter` (`updated_at`);--> statement-breakpoint
CREATE INDEX `idx_magazine_name` ON `magazine` (`name`);--> statement-breakpoint
CREATE INDEX `idx_magazine_demo` ON `magazine` (`demo`);--> statement-breakpoint
CREATE INDEX `idx_magazine_manga_magazine_id` ON `magazine_manga` (`magazine_id`);--> statement-breakpoint
CREATE INDEX `idx_magazine_manga_manga_id` ON `magazine_manga` (`manga_id`);--> statement-breakpoint
CREATE INDEX `idx_manga_name` ON `manga` (`name`);--> statement-breakpoint
CREATE INDEX `idx_manga_demo` ON `manga` (`demo`);--> statement-breakpoint
CREATE INDEX `idx_manga_genre_manga_id` ON `manga_genre` (`manga_id`);--> statement-breakpoint
CREATE INDEX `idx_manga_genre_genre` ON `manga_genre` (`genre`);--> statement-breakpoint
CREATE UNIQUE INDEX `unique_manga_genre` ON `manga_genre` (`manga_id`,`genre`);--> statement-breakpoint
CREATE INDEX `idx_manga_mangaka_job_manga_id` ON `manga_mangaka_job` (`manga_id`);--> statement-breakpoint
CREATE INDEX `idx_manga_mangaka_job_mangaka_id` ON `manga_mangaka_job` (`mangaka_id`);--> statement-breakpoint
CREATE UNIQUE INDEX `unique_manga_mangaka` ON `manga_mangaka_job` (`manga_id`,`mangaka_id`);--> statement-breakpoint
CREATE INDEX `idx_mangaka_name` ON `mangaka` (`name`);