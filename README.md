Naturebyte Forager Tool

Naturebyte Forager Tool is a command-line utility designed for foraging subdomains and running various reconnaissance tools on a given domain. This tool integrates with subfinder, theHarvester, dnsrecon, and whois to provide comprehensive domain reconnaissance.

Features:

Subdomain Discovery: Uses subfinder to discover subdomains of a given domain.
Reconnaissance Tools: Runs theHarvester, dnsrecon, and whois for further reconnaissance on the domain.

Installation

Follow these steps to install the Naturebyte Forager Tool on your system.
======================================================================================
Step 1: Clone the Repository

First, clone the repository from GitHub to your local machine:

git clone https://github.com/grimmtar/naturebyte.git
cd naturebyte
=====================================================================================
Step 2: Make the Installation Script Executable

Change the permissions of the installation script to make it executable:

chmod +x install_forage.sh
=====================================================================================
Step 3: Run the Installation Script

Execute the installation script to install the Forager tool and its dependencies:

./install_forage.sh
=====================================================================================
Step 4: Verify Installation

After running the installation script, verify that the tool is installed correctly by running:

forage example.com

You should see the tool executing and foraging subdomains for the given domain.

=======================================================================================

Troubleshooting

If you encounter any issues during installation or usage, please check the following:

Ensure you have internet connectivity to clone the repository and install dependencies.
Ensure you have the necessary permissions to execute the installation script and move the binary to /usr/local/bin.
If the problem persists, feel free to open an issue on the GitHub repository.

Contributing
We welcome contributions! If you'd like to contribute to this project, please fork the repository and submit a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for details.
