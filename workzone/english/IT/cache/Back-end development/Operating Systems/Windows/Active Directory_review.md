1) Q: What is the default Kerberos ticket-granting ticket (TGT) lifetime in Active Directory, and can it be modified? 
A: The default TGT lifetime is 10 hours, and it can be modified using the MaxTicketAge attribute in Active Directory.

2) Q: How can you manually force replication between two domain controllers in Active Directory? 
A: You can use the Repadmin.exe command-line tool, specifically the /syncall switch, to manually force replication between two domain controllers.

3) Q: What is the purpose of the Group Policy Central Store in Active Directory? 
A: The Group Policy Central Store provides a standardized location for storing group policy templates, ensuring that all administrators are using the same templates for managing group policies.

4) Q: What is the difference between the NTLM and Kerberos authentication protocols used in Active Directory? 
A: NTLM relies on a challenge-response mechanism and is less secure than Kerberos, which uses mutual authentication and encryption to protect against attacks such as replay attacks.

5) Q: How can you determine which domain controller a client machine has authenticated with in Active Directory? 
A: You can use the "set" command in Command Prompt on the client machine to view the LOGONSERVER environment variable, which will display the name of the domain controller the machine is authenticated with.