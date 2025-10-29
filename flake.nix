{
  description = "TaskHerald - Taskwarrior notification service";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    home-manager.url = "github:nix-community/home-manager";
  };

  outputs = { self, nixpkgs, flake-utils, home-manager }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages = {
          taskherald = pkgs.buildGoModule {
            pname = "taskherald";
            version = "0.1.0";
            src = ./.;

            vendorHash = "sha256-DD0C5oV44BNAKxStrVYm8KhttUSoJBAdvZlyuxq6Cqs=";

            subPackages = [ "src" ];

            postInstall = ''
              mv $out/bin/src $out/bin/taskherald
            '';

            meta = {
              description = "Taskwarrior notification service";
              license = pkgs.lib.licenses.mit;
              maintainers = [ ];
              mainProgram = "taskherald";
            };
          };
        };

        defaultPackage = self.packages.${system}.taskherald;
      }) // {
        homeManagerModules = {
          taskherald = { config, lib, pkgs, ... }: {
            options.services.taskherald = {
              enable = lib.mkEnableOption "TaskHerald notification service";

              settings = {
                ntfy_topic = lib.mkOption {
                  type = lib.types.str;
                  description = "Ntfy topic for notifications";
                };

                ntfy_server = lib.mkOption {
                  type = lib.types.str;
                  default = "https://ntfy.sh";
                  description = "Ntfy server URL";
                };

                taskherald_interval = lib.mkOption {
                  type = lib.types.int;
                  default = 60;
                  description = "Check interval in seconds";
                };
              };
            };

            config = lib.mkIf config.services.taskherald.enable {
              systemd.user.services.taskherald = {
                Unit = {
                  Description = "TaskHerald notification service";
                  After = [ "network.target" ];
                };

                Service = {
                  Type = "simple";
                  ExecStart = "${self.packages.${pkgs.system}.taskherald}/bin/taskherald";
                  Environment = [
                    "NTFY_TOPIC=${config.services.taskherald.settings.ntfy_topic}"
                    "NTFY_SERVER=${config.services.taskherald.settings.ntfy_server}"
                    "TASKHERALD_INTERVAL=${toString config.services.taskherald.settings.taskherald_interval}"
                  ];
                  Restart = "always";
                };

                Install = {
                  WantedBy = [ "default.target" ];
                };
              };
            };
          };
        };
      };
}