import pandas as pd
import matplotlib.pyplot as plt

rest = pd.read_csv("rest_latencies.csv")
grpc = pd.read_csv("grpc_latencies.csv")

plt.figure()
plt.hist(rest["latency_ms"], bins=50)
plt.title("REST Latency Distribution")
plt.xlabel("Latency (ms)")
plt.ylabel("Frequency")
plt.savefig("rest_hist.png")

plt.figure()
plt.hist(grpc["latency_ms"], bins=50)
plt.title("gRPC Latency Distribution")
plt.xlabel("Latency (ms)")
plt.ylabel("Frequency")
plt.savefig("grpc_hist.png")

# CDF plot
plt.figure()

for data, label in [(rest, "REST"), (grpc, "gRPC")]:
    sorted_data = sorted(data["latency_ms"])
    y = [i / len(sorted_data) for i in range(len(sorted_data))]
    plt.plot(sorted_data, y, label=label)

plt.title("CDF of Latency")
plt.xlabel("Latency (ms)")
plt.ylabel("CDF")
plt.legend()
plt.savefig("cdf.png")

plt.show()