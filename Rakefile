#
# Rakefile for agent-payload
#

protoc_binary="protoc"
protoc_version="3.5.1"
gogo_dir="/tmp/gogo"

namespace :codegen do

  task :install_protoc do
    if `bash -c "protoc --version"` != "libprotoc ${protoc_version}"
      protoc_binary="/tmp/protoc#{protoc_version}"
      sh <<-EOF
        /bin/bash <<BASH
        if [ ! -f #{protoc_binary} ] ; then
          echo "Downloading protoc #{protoc_version}"
          cd /tmp
          if [ "$(uname -s)" = "Darwin" ] ; then
            curl -OL https://github.com/google/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-osx-x86_64.zip
          else
            curl -OL https://github.com/google/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-linux-x86_64.zip
          fi
          unzip protoc-#{protoc_version}*.zip
          mv bin/protoc #{protoc_binary}
        fi
BASH
      EOF
    end
  end

  task :protoc => ['install_protoc'] do
    sh <<-EOF
      /bin/bash <<BASH
      set -euo pipefail

      export GO111MODULE=auto

      rm -rf #{gogo_dir}
      rm -rf /tmp/gogo-bin-*

      mkdir -p #{gogo_dir}/src/github.com/gogo
      git clone https://github.com/gogo/protobuf.git #{gogo_dir}/src/github.com/gogo/protobuf

      # Install v1.0.0
      pushd #{gogo_dir}/src/github.com/gogo/protobuf
      git checkout v1.0.0
      GOBIN=/tmp/gogo-bin-v1.0.0 GOPATH=#{gogo_dir} make clean install

      popd

      echo "Generating logs proto"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src --java_out=java proto/logs/agent_logs_payload.proto

      echo "Generating metrics proto (go)"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/metrics/agent_payload.proto

      echo "Generating metrics proto (python)"
      #{protoc_binary} --proto_path=#{gogo_dir}/src:$GOPATH/src:./proto/metrics --python_out=python agent_payload.proto

      # Install the specific tag that the process-agent needs
      pushd #{gogo_dir}/src/github.com/gogo/protobuf
      git checkout d76fbc1373015ced59b43ac267f28d546b955683
      GOBIN=/tmp/gogo-bin-d76fbc1373015ced59b43ac267f28d546b955683 GOPATH=#{gogo_dir} make clean install

      popd

      echo "Generating process proto"
      PATH=/tmp/gogo-bin-d76fbc1373015ced59b43ac267f28d546b955683 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofaster_out=$GOPATH/src proto/process/*.proto

      echo "Generating contlcycle proto"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/contlcycle/contlcycle.proto

      cp -r v5/* .
      rm -rf v5
BASH
    EOF
  end

  desc 'Run all code generators.'
  multitask :all => [:protoc]

end

# no desc, doesn't really work outside of rake tasks, can't say: `rake gimme; go` as the vars aren't exported
task :gimme do
    go_version = "1.18"

    if (`which gimme`; $?.success?)
      sh "GIMME_DOWNLOAD_BASE=https://binaries.ddbuild.io/golang/ gimme #{go_version}"
      results = `. ~/.gimme/envs/go#{go_version}.env > /dev/null; echo $GOOS; echo $GOARCH; echo $GOROOT; echo $PATH`.split("\n")
      ENV['GOOS'] = results[0]
      ENV['GOARCH'] = results[1]
      ENV['GOROOT'] = results[2]
      ENV['PATH'] = results[3]
    elsif (`which go`; $?.success?)
      puts "using local go toolchain, install gimme to use pinned go version"
    else
      puts "You need either gimme or go in your path."
      exit 1

  end
end

desc "Setup dependencies"
task :deps do
  system("go mod tidy")
end

desc "Run tests"
task :test do
  cmd = "go list ./... | grep -v vendor | xargs go test -v "
  sh cmd
end

task :fuzz do
  cmd = "go list ./... | grep -v vendor | xargs -n 1 go test -v -fuzztime=20s -fuzz '.*' "
  sh cmd
end

desc "Run all code generation."
task :codegen => ['codegen:all']

desc "Run all protobuf code generation."
task :protobuf => ['codegen:protoc']

task :default => [:gimme, :deps, :test, :codegen]
