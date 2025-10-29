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
                  type = lib.types.nullOr lib.types.str;
                  default = null;
                  description = "Ntfy topic for notifications (mutually exclusive with ntfy_topic_file)";
                };

                ntfy_topic_file = lib.mkOption {
                  type = lib.types.nullOr lib.types.path;
                  default = null;
                  description = "Path to file containing ntfy topic (mutually exclusive with ntfy_topic)";
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
              assertions = [
                {
                  assertion = (config.services.taskherald.settings.ntfy_topic != null) != (config.services.taskherald.settings.ntfy_topic_file != null);
                  message = "Exactly one of services.taskherald.settings.ntfy_topic or services.taskherald.settings.ntfy_topic_file must be set";
                }
              ];

              systemd.user.services.taskherald = {
                Unit = {
                  Description = "TaskHerald notification service";
                  After = [ "network.target" ];
                };

                Service = {
                  Type = "simple";
                  ExecStart = "${self.packages.${pkgs.system}.taskherald}/bin/taskherald";
                  Environment = lib.mkMerge [
                    [
                      "NTFY_SERVER=${config.services.taskherald.settings.ntfy_server}"
                      "TASKHERALD_INTERVAL=${toString config.services.taskherald.settings.taskherald_interval}"
                    ]
                    (lib.mkIf (config.services.taskherald.settings.ntfy_topic != null) [
                      "NTFY_TOPIC=${config.services.taskherald.settings.ntfy_topic}"
                    ])
                    (lib.mkIf (config.services.taskherald.settings.ntfy_topic_file != null) [
                      "NTFY_TOPIC_FILE=${config.services.taskherald.settings.ntfy_topic_file}"
                    ])
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