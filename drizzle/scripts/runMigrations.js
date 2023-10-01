import { execSync } from "node:child_process";
import { existsSync, readdirSync, unlinkSync } from "node:fs";
import { join } from "node:path";
import { fileURLToPath } from "node:url";

// there is also a --command="SQL STRING" option for npx wrangler d1 you can use for one off commands like queries

const FILE_SYSTEM_DB_FILE_NAME = "no-manga.db";
const CLEAR_FILE_NAME = "clear.sql";
const SEED_FILE_NAME = "seed.sql";

const SUPPRESS_D1_WARNING = "export NO_D1_WARNING=true";

const filesPath = join(fileURLToPath(import.meta.url), "..", "..");

const migrationFiles = [
  CLEAR_FILE_NAME,
  ...readdirSync(filesPath)
    .filter(
      (file) =>
        file.endsWith(".sql") &&
        file !== CLEAR_FILE_NAME &&
        file !== SEED_FILE_NAME,
    )
    .sort(),
  SEED_FILE_NAME,
].map((file) => join(filesPath, file));

try {
  const FsDb = join(filesPath, FILE_SYSTEM_DB_FILE_NAME);
  if (existsSync(FsDb)) unlinkSync(FsDb);

  console.log(">>>FILE_SYSTEM");
  migrationFiles.forEach((migFile) => {
    console.log(migFile);
    execSync(`cat ${migFile} | sqlite3 ${FsDb}`);
  });
} catch (e) {
  console.error(e);
}

try {
  console.log(">>>LOCAL WRANGLER");
  migrationFiles.forEach((migFile) => {
    console.log(migFile);
    execSync(
      `${SUPPRESS_D1_WARNING} && npx wrangler d1 execute no-manga --local --file=${migFile}`,
    );
  });

  console.log(">>>REMOTE WRANGLER");
  migrationFiles.forEach((migFile) => {
    console.log(migFile);
    execSync(
      `${SUPPRESS_D1_WARNING} && npx wrangler d1 execute no-manga --file=${migFile}`,
    );
  });
} catch (e) {
  // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-call
  console.error(e.stdout?.toString?.());
  // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-call
  console.error(e.stderr?.toString?.());
}
