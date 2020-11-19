import requests
from bs4 import BeautifulSoup


def get_proxies() -> list:
    proxies = []
    response = requests.get("https://free-proxy-list.net/")
    soup = BeautifulSoup(response.text, "html.parser")

    for tr in soup.find_all("tr"):
        td_list = tr.find_all("td")
        proxy_dict = {}
        if td_list is not None and len(td_list) == 8:
            for n, key in enumerate(
                [
                    "IP",
                    "Port",
                    "code",
                    "country",
                    "anonymity",
                    "google",
                    "https",
                    "refresh",
                ]
            ):
                text = td_list[n].get_text()
                if text == "yes":
                    text = True
                elif text == "no":
                    text = False
                proxy_dict[key] = text
            proxies.append(proxy_dict)
    return proxies


if __name__ == "__main__":
    for proxy in get_proxies():
        if proxy.get("https"):
            print(proxy.get("IP") + ":" + proxy.get("Port"))
