FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build-env
WORKDIR /App

# Disable ipv6
RUN echo "net.ipv6.conf.all.disable_ipv6 = 1" >> /etc/sysctl.conf
RUN echo "net.ipv6.conf.default.disable_ipv6 = 1" >> /etc/sysctl.conf 

# Copy everything
COPY . ./

# Restore as distinct layers
RUN dotnet restore
RUN dotnet build
RUN dotnet test


# Build and publish a release
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /App
COPY --from=build-env /App/out .

ENV ASPNETCORE_URLS=http://localhost:80

EXPOSE 80
ENTRYPOINT ["dotnet", "CoproXpert.dll"]

